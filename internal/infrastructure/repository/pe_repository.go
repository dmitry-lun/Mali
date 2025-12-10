package repository

import (
	"github.com/dmitry-lun/Mali/internal/domain/entity"
	"github.com/dmitry-lun/Mali/internal/domain/repository"
	"github.com/dmitry-lun/Mali/internal/infrastructure/parser"
	"github.com/dmitry-lun/Mali/pkg/detector"
	"github.com/dmitry-lun/Mali/pkg/file"
	"github.com/dmitry-lun/Mali/pkg/hash"
)

type peRepository struct {
	fileReader file.Reader
	hasher     hash.Hasher
}

func NewPERepository(fileReader file.Reader) repository.PERepository {
	return &peRepository{
		fileReader: fileReader,
		hasher:     hash.NewFileHasher(),
	}
}

func (r *peRepository) Parse(path string) (*entity.PEFile, error) {
	data, err := r.fileReader.ReadFile(path)
	if err != nil {
		return nil, err
	}

	isPE, lfanew, err := parser.ValidatePE(data)
	if err != nil || !isPE {
		return nil, err
	}

	dosHeader, err := parser.ParseDOSHeader(data)
	if err != nil {
		return nil, err
	}

	fileHeader, err := parser.ParseFileHeader(data, lfanew)
	if err != nil {
		return nil, err
	}

	optionalHeader, err := parser.ParseOptionalHeader(data, lfanew, fileHeader.SizeOfOptionalHeader)
	if err != nil {
		return nil, err
	}

	sections, err := parser.ParseSections(data, lfanew)
	if err != nil {
		return nil, err
	}

	importTableRVA, err := parser.GetImportTableRVA(data, lfanew, fileHeader.SizeOfOptionalHeader)
	if err != nil {
		importTableRVA = 0
	}

	var imports []entity.Import
	if importTableRVA != 0 {
		imports, err = parser.ParseImports(data, importTableRVA, sections)
		if err != nil {
			imports = []entity.Import{}
		}
	}

	exportTableRVA, err := parser.GetExportTableRVA(data, lfanew, fileHeader.SizeOfOptionalHeader)
	if err != nil {
		exportTableRVA = 0
	}

	var exports []entity.Export
	if exportTableRVA != 0 {
		exports, err = parser.ParseExports(data, exportTableRVA, sections)
		if err != nil {
			exports = []entity.Export{}
		}
	}

	packers := parser.DetectPackers(sections)

	machineStr := "unknown"
	if machineName, ok := validMachines[fileHeader.Machine]; ok {
		machineStr = machineName
	}

	entropy := make(map[string]float64)
	for _, section := range sections {
		if int(section.PointerToRawData)+int(section.RawSize) <= len(data) {
			sectionData := data[section.PointerToRawData : section.PointerToRawData+section.RawSize]
			entropy[section.Name] = parser.CalculateEntropy(sectionData)
		}
	}

	md5Hash := r.hasher.HashMD5(data)
	sha256Hash := r.hasher.HashSHA256(data)

	peFile := &entity.PEFile{
		Path:       path,
		IsPE:       isPE,
		DOSHeader:  dosHeader,
		FileHeader: fileHeader,
		Optional:   optionalHeader,
		Sections:   sections,
		Imports:    imports,
		Exports:    exports,
		EntryPoint: optionalHeader.AddressOfEntryPoint,
		Machine:    machineStr,
		Entropy:    entropy,
		Packers:    packers,
		Size:       uint64(len(data)),
		Resources:  []entity.Resource{},
		MD5:        md5Hash,
		SHA256:     sha256Hash,
	}

	peFile.RiskLevel = detector.CalculateRiskLevel(peFile)

	return peFile, nil
}

var validMachines = map[uint16]string{
	0x014c: "x86",
	0x8664: "x64",
	0x01c0: "ARM",
	0x01c4: "ARMv7",
	0xaa64: "ARM64",
}
