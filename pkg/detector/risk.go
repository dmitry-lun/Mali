package detector

import "github.com/dmitry-lun/Mali/internal/domain/entity"

func CalculateRiskLevel(peFile *entity.PEFile) string {
	riskScore := 0

	if len(peFile.Packers) > 0 {
		riskScore += 30
	}

	highEntropyCount := 0
	for _, entropy := range peFile.Entropy {
		if entropy > 7.0 {
			highEntropyCount++
		}
	}
	if highEntropyCount > 0 {
		riskScore += highEntropyCount * 10
	}

	suspiciousImports := []string{
		"VirtualAlloc", "VirtualProtect", "WriteProcessMemory",
		"CreateRemoteThread", "NtCreateThreadEx", "NtAllocateVirtualMemory",
		"GetProcAddress", "LoadLibrary", "CreateFileA", "WriteFile",
	}

	hasSuspiciousImport := false
	for _, imp := range peFile.Imports {
		for _, funcName := range imp.Functions {
			for _, suspicious := range suspiciousImports {
				if funcName == suspicious {
					hasSuspiciousImport = true
					riskScore += 15
					break
				}
			}
			if hasSuspiciousImport {
				break
			}
		}
		if hasSuspiciousImport {
			break
		}
	}

	if len(peFile.Sections) == 0 {
		riskScore += 20
	}

	if riskScore >= 60 {
		return "HIGH"
	} else if riskScore >= 30 {
		return "MEDIUM"
	} else if riskScore > 0 {
		return "LOW"
	}
	return "SAFE"
}
