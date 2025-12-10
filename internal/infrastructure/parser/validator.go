package parser

import (
	"encoding/binary"
	"fmt"
)

var validMachines = map[uint16]string{
	0x014c: "x86",
	0x8664: "x64",
	0x01c0: "ARM",
	0x01c4: "ARMv7",
	0xaa64: "ARM64",
}

func checkSize(data []byte) error {
	if len(data) < 0x40 {
		return fmt.Errorf("file too small to be PE")
	}
	return nil
}

func checkDOSHeader(data []byte) (uint32, error) {
	if data[0] != 'M' || data[1] != 'Z' {
		return 0, fmt.Errorf("file is not PE (missing MZ)")
	}

	lfanew := binary.LittleEndian.Uint32(data[0x3C:0x40])
	if int(lfanew)+24 > len(data) {
		return 0, fmt.Errorf("invalid lfanew (PE header offset)")
	}

	return lfanew, nil
}

func checkPEHeader(data []byte, lfanew uint32) error {
	if data[lfanew] != 'P' || data[lfanew+1] != 'E' || data[lfanew+2] != 0 || data[lfanew+3] != 0 {
		return fmt.Errorf("missing PE header")
	}
	return nil
}

func checkSectionCount(data []byte, lfanew uint32) error {
	fileHeader := lfanew + 4

	machine := binary.LittleEndian.Uint16(data[fileHeader : fileHeader+2])
	numberOfSections := binary.LittleEndian.Uint16(data[fileHeader+2 : fileHeader+4])

	if _, ok := validMachines[machine]; !ok {
		return fmt.Errorf("invalid or unknown Machine type: 0x%x", machine)
	}

	if numberOfSections == 0 || numberOfSections > 100 {
		return fmt.Errorf("invalid NumberOfSections: %d", numberOfSections)
	}

	return nil
}

func ValidatePE(data []byte) (bool, uint32, error) {
	if err := checkSize(data); err != nil {
		return false, 0, err
	}

	lfanew, err := checkDOSHeader(data)
	if err != nil {
		return false, 0, err
	}

	if err := checkPEHeader(data, lfanew); err != nil {
		return false, 0, err
	}

	if err := checkSectionCount(data, lfanew); err != nil {
		return false, 0, err
	}

	return true, lfanew, nil
}
