package binary

import "fmt"

type Section interface {
	GetVirtualAddress() uint32
	GetVirtualSize() uint32
	GetPointerToRawData() uint32
}

func RVAtoRaw(rva uint32, sections []Section) (uint32, error) {
	for _, sec := range sections {
		if rva >= sec.GetVirtualAddress() && rva < sec.GetVirtualAddress()+sec.GetVirtualSize() {
			offset := rva - sec.GetVirtualAddress() + sec.GetPointerToRawData()
			return offset, nil
		}
	}
	return 0, fmt.Errorf("RVA 0x%x does not fall into any section", rva)
}
