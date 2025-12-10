package entity

type PEFile struct {
	Path       string             `json:"path"`
	IsPE       bool               `json:"isPE"`
	Machine    string             `json:"machine"`
	EntryPoint uint32             `json:"entryPoint"`
	Size       uint64             `json:"size"`
	MD5        string             `json:"md5"`
	SHA256     string             `json:"sha256"`
	RiskLevel  string             `json:"riskLevel"`
	Packers    []string           `json:"packers,omitempty"`
	Entropy    map[string]float64 `json:"entropy,omitempty"`
	Sections   []Section          `json:"sections,omitempty"`
	Imports    []Import           `json:"imports,omitempty"`
	Exports    []Export           `json:"exports,omitempty"`
	DOSHeader  DOSHeader          `json:"dosHeader,omitempty"`
	FileHeader FileHeader         `json:"fileHeader,omitempty"`
	Optional   OptionalHeader     `json:"optionalHeader,omitempty"`
	Resources  []Resource         `json:"resources,omitempty"`
}

type Section struct {
	Name             string `json:"name"`
	VirtualAddress   uint32 `json:"virtualAddress"`
	VirtualSize      uint32 `json:"virtualSize"`
	RawSize          uint32 `json:"rawSize"`
	PointerToRawData uint32 `json:"pointerToRawData"`
	Characteristics  uint32 `json:"characteristics"`
	RawData          []byte `json:"rawData,omitempty"`
}

func (s Section) GetName() string {
	return s.Name
}

func (s Section) GetVirtualAddress() uint32 {
	return s.VirtualAddress
}

func (s Section) GetVirtualSize() uint32 {
	return s.VirtualSize
}

func (s Section) GetPointerToRawData() uint32 {
	return s.PointerToRawData
}

type Import struct {
	DLL       string   `json:"dll"`
	Functions []string `json:"functions"`
}

type Export struct {
	Name    string `json:"name"`
	Address uint32 `json:"address"`
}

type Resource struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Data []byte `json:"data,omitempty"`
}
