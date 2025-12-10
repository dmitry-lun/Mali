package parser

import (
	"encoding/binary"

	"github.com/dmitry-lun/Mali/internal/domain/entity"
)

func ParseExports(data []byte, exportTableRVA uint32, sections []entity.Section) ([]entity.Export, error) {
	offset, err := RVAtoRaw(exportTableRVA, sections)
	if err != nil {
		return nil, err
	}

	var exports []entity.Export
	if int(offset)+40 <= len(data) {
		namesRVA := binary.LittleEndian.Uint32(data[offset+24 : offset+28])
		namesCount := binary.LittleEndian.Uint32(data[offset+20 : offset+24])

		for i := uint32(0); i < namesCount; i++ {
			nameRVA := binary.LittleEndian.Uint32(data[offset+namesRVA+i*4 : offset+namesRVA+(i+1)*4])
			nameOffset, err := RVAtoRaw(nameRVA, sections)
			if err != nil {
				return nil, err
			}
			name := ""
			for j := nameOffset; j < uint32(len(data)) && data[j] != 0; j++ {
				name += string(data[j])
			}
			exports = append(exports, entity.Export{Name: name})
		}
	}

	return exports, nil
}
