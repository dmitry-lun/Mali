package parser

import (
	"encoding/binary"
	"fmt"

	"github.com/dmitry-lun/Mali/internal/domain/entity"
	binaryutil "github.com/dmitry-lun/Mali/pkg/binary"
)

func ParseSections(data []byte, lfanew uint32) ([]entity.Section, error) {
	fileHeaderOffset := lfanew + 4
	numberOfSections := binary.LittleEndian.Uint16(data[fileHeaderOffset+2 : fileHeaderOffset+4])
	sizeOfOptionalHeader := binary.LittleEndian.Uint16(data[fileHeaderOffset+16 : fileHeaderOffset+18])
	sectionTableOffset := lfanew + 24 + uint32(sizeOfOptionalHeader)

	sections := make([]entity.Section, 0, numberOfSections)

	for i := uint16(0); i < numberOfSections; i++ {
		start := sectionTableOffset + uint32(i*40)

		nameBytes := data[start : start+8]
		name := string(nameBytes)
		for j, b := range nameBytes {
			if b == 0 {
				name = string(nameBytes[:j])
				break
			}
		}

		virtualSize := binary.LittleEndian.Uint32(data[start+8 : start+12])
		virtualAddress := binary.LittleEndian.Uint32(data[start+12 : start+16])
		sizeOfRawData := binary.LittleEndian.Uint32(data[start+16 : start+20])
		pointerToRawData := binary.LittleEndian.Uint32(data[start+20 : start+24])
		characteristics := binary.LittleEndian.Uint32(data[start+36 : start+40])

		sections = append(sections, entity.Section{
			Name:             name,
			VirtualAddress:   virtualAddress,
			VirtualSize:      virtualSize,
			RawSize:          sizeOfRawData,
			PointerToRawData: pointerToRawData,
			Characteristics:  characteristics,
		})
	}

	if len(sections) == 0 {
		return nil, fmt.Errorf("no sections found")
	}

	return sections, nil
}

func RVAtoRaw(rva uint32, sections []entity.Section) (uint32, error) {
	sectionInterfaces := make([]binaryutil.Section, len(sections))
	for i := range sections {
		sectionInterfaces[i] = sections[i]
	}
	return binaryutil.RVAtoRaw(rva, sectionInterfaces)
}
