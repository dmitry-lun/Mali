package entity

type DOSHeader struct {
	Magic  [2]byte `json:"magic"`
	Lfanew uint32  `json:"lfanew"`
}

type FileHeader struct {
	Machine              uint16 `json:"machine"`
	NumberOfSections     uint16 `json:"numberOfSections"`
	TimeDateStamp        uint32 `json:"timeDateStamp"`
	PointerToSymbolTable uint32 `json:"pointerToSymbolTable"`
	NumberOfSymbols      uint32 `json:"numberOfSymbols"`
	SizeOfOptionalHeader uint16 `json:"sizeOfOptionalHeader"`
	Characteristics      uint16 `json:"characteristics"`
}

type OptionalHeader struct {
	Magic                   uint16 `json:"magic"`
	MajorLinkerVersion      byte   `json:"majorLinkerVersion"`
	MinorLinkerVersion      byte   `json:"minorLinkerVersion"`
	SizeOfCode              uint32 `json:"sizeOfCode"`
	SizeOfInitializedData   uint32 `json:"sizeOfInitializedData"`
	SizeOfUninitializedData uint32 `json:"sizeOfUninitializedData"`
	AddressOfEntryPoint     uint32 `json:"addressOfEntryPoint"`
	BaseOfCode              uint32 `json:"baseOfCode"`
	ImageBase               uint64 `json:"imageBase"`
	SectionAlignment        uint32 `json:"sectionAlignment"`
	FileAlignment           uint32 `json:"fileAlignment"`
	SizeOfImage             uint32 `json:"sizeOfImage"`
	SizeOfHeaders           uint32 `json:"sizeOfHeaders"`
	Subsystem               uint16 `json:"subsystem"`
	DllCharacteristics      uint16 `json:"dllCharacteristics"`
	SizeOfStackReserve      uint64 `json:"sizeOfStackReserve"`
	SizeOfHeapReserve       uint64 `json:"sizeOfHeapReserve"`
}

type SectionHeader struct {
	Name             [8]byte
	VirtualSize      uint32
	VirtualAddress   uint32
	SizeOfRawData    uint32
	PointerToRawData uint32
	PointerToReloc   uint32
	PointerToLinenum uint32
	NumReloc         uint16
	NumLineNumbers   uint16
	Characteristics  uint32
}
