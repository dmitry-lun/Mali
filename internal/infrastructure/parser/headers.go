package parser

import (
	"encoding/binary"
	"fmt"

	"github.com/dmitry-lun/Mali/internal/domain/entity"
)

func ParseDOSHeader(data []byte) (entity.DOSHeader, error) {
	if len(data) < 0x40 {
		return entity.DOSHeader{}, fmt.Errorf("file too small for DOS header")
	}

	var dosHeader entity.DOSHeader
	dosHeader.Magic[0] = data[0]
	dosHeader.Magic[1] = data[1]
	dosHeader.Lfanew = binary.LittleEndian.Uint32(data[0x3C:0x40])

	return dosHeader, nil
}

func ParseFileHeader(data []byte, lfanew uint32) (entity.FileHeader, error) {
	fileHeaderOffset := lfanew + 4
	if int(fileHeaderOffset)+20 > len(data) {
		return entity.FileHeader{}, fmt.Errorf("file too small for File Header")
	}

	var fileHeader entity.FileHeader
	fileHeader.Machine = binary.LittleEndian.Uint16(data[fileHeaderOffset : fileHeaderOffset+2])
	fileHeader.NumberOfSections = binary.LittleEndian.Uint16(data[fileHeaderOffset+2 : fileHeaderOffset+4])
	fileHeader.TimeDateStamp = binary.LittleEndian.Uint32(data[fileHeaderOffset+4 : fileHeaderOffset+8])
	fileHeader.PointerToSymbolTable = binary.LittleEndian.Uint32(data[fileHeaderOffset+8 : fileHeaderOffset+12])
	fileHeader.NumberOfSymbols = binary.LittleEndian.Uint32(data[fileHeaderOffset+12 : fileHeaderOffset+16])
	fileHeader.SizeOfOptionalHeader = binary.LittleEndian.Uint16(data[fileHeaderOffset+16 : fileHeaderOffset+18])
	fileHeader.Characteristics = binary.LittleEndian.Uint16(data[fileHeaderOffset+18 : fileHeaderOffset+20])

	return fileHeader, nil
}

func ParseOptionalHeader(data []byte, lfanew uint32, sizeOfOptionalHeader uint16) (entity.OptionalHeader, error) {
	optionalHeaderOffset := lfanew + 24
	if int(optionalHeaderOffset)+int(sizeOfOptionalHeader) > len(data) {
		return entity.OptionalHeader{}, fmt.Errorf("file too small for Optional Header")
	}

	var optionalHeader entity.OptionalHeader
	optionalHeader.Magic = binary.LittleEndian.Uint16(data[optionalHeaderOffset : optionalHeaderOffset+2])
	optionalHeader.MajorLinkerVersion = data[optionalHeaderOffset+2]
	optionalHeader.MinorLinkerVersion = data[optionalHeaderOffset+3]
	optionalHeader.SizeOfCode = binary.LittleEndian.Uint32(data[optionalHeaderOffset+4 : optionalHeaderOffset+8])
	optionalHeader.SizeOfInitializedData = binary.LittleEndian.Uint32(data[optionalHeaderOffset+8 : optionalHeaderOffset+12])
	optionalHeader.SizeOfUninitializedData = binary.LittleEndian.Uint32(data[optionalHeaderOffset+12 : optionalHeaderOffset+16])
	optionalHeader.AddressOfEntryPoint = binary.LittleEndian.Uint32(data[optionalHeaderOffset+16 : optionalHeaderOffset+20])
	optionalHeader.BaseOfCode = binary.LittleEndian.Uint32(data[optionalHeaderOffset+20 : optionalHeaderOffset+24])

	if optionalHeader.Magic == 0x10b {
		optionalHeader.ImageBase = uint64(binary.LittleEndian.Uint32(data[optionalHeaderOffset+28 : optionalHeaderOffset+32]))
		optionalHeader.SectionAlignment = binary.LittleEndian.Uint32(data[optionalHeaderOffset+32 : optionalHeaderOffset+36])
		optionalHeader.FileAlignment = binary.LittleEndian.Uint32(data[optionalHeaderOffset+36 : optionalHeaderOffset+40])
		optionalHeader.SizeOfImage = binary.LittleEndian.Uint32(data[optionalHeaderOffset+56 : optionalHeaderOffset+60])
		optionalHeader.SizeOfHeaders = binary.LittleEndian.Uint32(data[optionalHeaderOffset+60 : optionalHeaderOffset+64])
		optionalHeader.Subsystem = binary.LittleEndian.Uint16(data[optionalHeaderOffset+68 : optionalHeaderOffset+70])
		optionalHeader.DllCharacteristics = binary.LittleEndian.Uint16(data[optionalHeaderOffset+70 : optionalHeaderOffset+72])
		optionalHeader.SizeOfStackReserve = uint64(binary.LittleEndian.Uint32(data[optionalHeaderOffset+72 : optionalHeaderOffset+76]))
		optionalHeader.SizeOfHeapReserve = uint64(binary.LittleEndian.Uint32(data[optionalHeaderOffset+76 : optionalHeaderOffset+80]))
	} else if optionalHeader.Magic == 0x20b {
		optionalHeader.ImageBase = binary.LittleEndian.Uint64(data[optionalHeaderOffset+24 : optionalHeaderOffset+32])
		optionalHeader.SectionAlignment = binary.LittleEndian.Uint32(data[optionalHeaderOffset+32 : optionalHeaderOffset+36])
		optionalHeader.FileAlignment = binary.LittleEndian.Uint32(data[optionalHeaderOffset+36 : optionalHeaderOffset+40])
		optionalHeader.SizeOfImage = binary.LittleEndian.Uint32(data[optionalHeaderOffset+56 : optionalHeaderOffset+60])
		optionalHeader.SizeOfHeaders = binary.LittleEndian.Uint32(data[optionalHeaderOffset+60 : optionalHeaderOffset+64])
		optionalHeader.Subsystem = binary.LittleEndian.Uint16(data[optionalHeaderOffset+68 : optionalHeaderOffset+70])
		optionalHeader.DllCharacteristics = binary.LittleEndian.Uint16(data[optionalHeaderOffset+70 : optionalHeaderOffset+72])
		optionalHeader.SizeOfStackReserve = binary.LittleEndian.Uint64(data[optionalHeaderOffset+80 : optionalHeaderOffset+88])
		optionalHeader.SizeOfHeapReserve = binary.LittleEndian.Uint64(data[optionalHeaderOffset+88 : optionalHeaderOffset+96])
	} else {
		return entity.OptionalHeader{}, fmt.Errorf("invalid optional header magic: 0x%x", optionalHeader.Magic)
	}

	return optionalHeader, nil
}

func GetImportTableRVA(data []byte, lfanew uint32, sizeOfOptionalHeader uint16) (uint32, error) {
	optionalHeaderOffset := lfanew + 24
	var importTableRVAOffset uint32

	if sizeOfOptionalHeader < 96 {
		return 0, fmt.Errorf("optional header too small for data directories")
	}

	magic := binary.LittleEndian.Uint16(data[optionalHeaderOffset : optionalHeaderOffset+2])
	if magic == 0x10b {
		importTableRVAOffset = optionalHeaderOffset + 104
	} else if magic == 0x20b {
		importTableRVAOffset = optionalHeaderOffset + 120
	} else {
		return 0, fmt.Errorf("invalid optional header magic: 0x%x", magic)
	}

	if int(importTableRVAOffset)+4 > len(data) {
		return 0, fmt.Errorf("file too small for Import Table RVA")
	}

	return binary.LittleEndian.Uint32(data[importTableRVAOffset : importTableRVAOffset+4]), nil
}

func GetExportTableRVA(data []byte, lfanew uint32, sizeOfOptionalHeader uint16) (uint32, error) {
	optionalHeaderOffset := lfanew + 24
	var exportTableRVAOffset uint32

	if sizeOfOptionalHeader < 96 {
		return 0, fmt.Errorf("optional header too small for data directories")
	}

	magic := binary.LittleEndian.Uint16(data[optionalHeaderOffset : optionalHeaderOffset+2])
	if magic == 0x10b {
		exportTableRVAOffset = optionalHeaderOffset + 96
	} else if magic == 0x20b {
		exportTableRVAOffset = optionalHeaderOffset + 112
	} else {
		return 0, fmt.Errorf("invalid optional header magic: 0x%x", magic)
	}

	if int(exportTableRVAOffset)+4 > len(data) {
		return 0, fmt.Errorf("file too small for Export Table RVA")
	}

	return binary.LittleEndian.Uint32(data[exportTableRVAOffset : exportTableRVAOffset+4]), nil
}
