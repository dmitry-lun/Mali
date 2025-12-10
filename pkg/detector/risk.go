package detector

import "github.com/dmitry-lun/Mali/internal/domain/entity"

type riskWeights struct {
	headerAnomalies     float64
	nonStandardSections float64
	entropy             float64
	importsExports      float64
	packers             float64
	hashDatabases       float64
}

var defaultWeights = riskWeights{
	headerAnomalies:     0.8,
	nonStandardSections: 0.7,
	entropy:             1.0,
	importsExports:      0.9,
	packers:             1.0,
	hashDatabases:       1.0,
}

func CalculateRiskLevel(peFile *entity.PEFile) string {
	weights := defaultWeights

	score := int(float64(scoreHeaderAnomalies(peFile))*weights.headerAnomalies) +
		int(float64(scoreNonStandardSections(peFile.Sections))*weights.nonStandardSections) +
		int(float64(scoreEntropy(peFile.Entropy))*weights.entropy) +
		int(float64(scoreImportsAndExports(peFile.Imports, peFile.Exports))*weights.importsExports) +
		int(float64(scorePackers(peFile.Packers))*weights.packers)

	switch {
	case score <= 5:
		return "SAFE"
	case score <= 10:
		return "LOW"
	case score <= 20:
		return "MEDIUM"
	default:
		return "HIGH"
	}
}

func scoreHeaderAnomalies(peFile *entity.PEFile) int {
	score := 0
	opt := peFile.Optional

	if peFile.DOSHeader.Magic[0] != 'M' || peFile.DOSHeader.Magic[1] != 'Z' {
		score += 5
	}
	if peFile.FileHeader.NumberOfSections == 0 {
		score += 5
	} else if peFile.FileHeader.NumberOfSections > 100 {
		score += 3
	}
	if opt.Magic != 0x10b && opt.Magic != 0x20b {
		score += 5
	}
	if opt.SectionAlignment == 0 || opt.FileAlignment == 0 || opt.SizeOfImage == 0 {
		score += 3
	}
	if opt.SectionAlignment > 0 && opt.SectionAlignment < 0x1000 {
		score += 2
	}
	if opt.FileAlignment > 0 && opt.FileAlignment < 0x200 && opt.FileAlignment != 0x1000 {
		score += 2
	}
	if opt.AddressOfEntryPoint == 0 {
		score += 2
	}

	return score
}

func scoreNonStandardSections(sections []entity.Section) int {
	standardSections := map[string]bool{
		".text": true, ".data": true, ".rdata": true, ".rsrc": true, ".reloc": true, ".bss": true,
		".edata": true, ".idata": true, ".pdata": true, ".xdata": true, ".tls": true, ".crt": true,
		".debug": true, ".drectve": true, ".didat": true, ".gfids": true, ".gljmp": true,
		".msvcjmp": true, ".sbss": true, ".sdata": true, ".srdata": true,
	}

	suspiciousSections := map[string]int{
		"UPX0": 5, "UPX1": 5, ".UPX": 5, ".themida": 5, ".vmp0": 5, ".vmp1": 5, ".aspack": 5,
		".text1": 4, ".data2": 4, ".data3": 4, ".packed": 4, ".pack": 4,
		".code": 3,
	}

	score := 0
	for _, sec := range sections {
		if points, found := suspiciousSections[sec.Name]; found {
			score += points
			continue
		}
		if !standardSections[sec.Name] && len(sec.Name) > 0 && sec.Name[0] == '.' {
			score += 2
		}
	}
	return score
}

func scoreEntropy(entropyMap map[string]float64) int {
	score := 0
	for _, e := range entropyMap {
		switch {
		case e >= 7.0:
			score += 10
		case e >= 4.5:
			score += 3
		}
	}
	return score
}

func scoreImportsAndExports(imports []entity.Import, exports []entity.Export) int {
	suspiciousAPIs := map[string]bool{
		"VirtualAlloc": true, "VirtualProtect": true, "WriteProcessMemory": true,
		"CreateRemoteThread": true, "NtCreateThreadEx": true, "NtAllocateVirtualMemory": true,
		"GetProcAddress": true, "LoadLibrary": true, "CreateFileA": true, "WriteFile": true,
		"NtWriteVirtualMemory": true, "NtProtectVirtualMemory": true, "SetWindowsHookEx": true,
		"CreateToolhelp32Snapshot": true, "Process32First": true, "Process32Next": true,
	}

	suspiciousCount := 0
	for _, imp := range imports {
		for _, fn := range imp.Functions {
			if suspiciousAPIs[fn] {
				suspiciousCount++
			}
		}
	}

	score := 0
	if suspiciousCount > 0 {
		switch {
		case suspiciousCount >= 5:
			score += 5
		case suspiciousCount >= 3:
			score += 3
		default:
			score += 1
		}
	}

	switch len(imports) {
	case 0:
		score += 5
	case 1, 2:
		score += 2
	}

	return score
}

func scorePackers(packers []string) int {
	if len(packers) == 0 {
		return 0
	}

	packerScores := map[string]int{
		"UPX": 5, "ASPack": 7, "Themida": 10, "VMProtect": 10,
	}

	maxScore := 0
	for _, p := range packers {
		if score, found := packerScores[p]; found && score > maxScore {
			maxScore = score
		} else if !found && maxScore < 8 {
			maxScore = 8
		}
	}
	return maxScore
}
