package parser

import (
	"encoding/binary"

	"github.com/dmitry-lun/Mali/internal/domain/entity"
)

func ParseImports(data []byte, importTableRVA uint32, sections []entity.Section) ([]entity.Import, error) {
	offset, err := RVAtoRaw(importTableRVA, sections)
	if err != nil {
		return nil, err
	}

	var imports []entity.Import

	for {
		if int(offset)+20 > len(data) {
			break
		}

		origFirstThunk := binary.LittleEndian.Uint32(data[offset : offset+4])
		timeDateStamp := binary.LittleEndian.Uint32(data[offset+4 : offset+8])
		forwarderChain := binary.LittleEndian.Uint32(data[offset+8 : offset+12])
		nameRVA := binary.LittleEndian.Uint32(data[offset+12 : offset+16])
		firstThunk := binary.LittleEndian.Uint32(data[offset+16 : offset+20])

		if origFirstThunk == 0 && timeDateStamp == 0 && forwarderChain == 0 && nameRVA == 0 && firstThunk == 0 {
			break
		}

		nameOffset, err := RVAtoRaw(nameRVA, sections)
		if err != nil {
			return nil, err
		}

		name := ""
		for i := nameOffset; i < uint32(len(data)) && data[i] != 0; i++ {
			name += string(data[i])
		}

		funcs, err := parseImportedFunctions(data, origFirstThunk, sections)
		if err != nil {
			return nil, err
		}

		imports = append(imports, entity.Import{
			DLL:       name,
			Functions: funcs,
		})

		offset += 20
	}

	return imports, nil
}

func parseImportedFunctions(data []byte, thunkRVA uint32, sections []entity.Section) ([]string, error) {
	offset, err := RVAtoRaw(thunkRVA, sections)
	if err != nil {
		return nil, err
	}

	var funcs []string

	for {
		if int(offset)+4 > len(data) {
			break
		}

		thunk := binary.LittleEndian.Uint32(data[offset : offset+4])
		if thunk == 0 {
			break
		}

		hintNameRVA := thunk & 0x7FFFFFFF
		hintNameOffset, err := RVAtoRaw(hintNameRVA, sections)
		if err != nil {
			return nil, err
		}

		name := ""
		for i := hintNameOffset + 2; i < uint32(len(data)) && data[i] != 0; i++ {
			name += string(data[i])
		}

		funcs = append(funcs, name)
		offset += 4
	}

	return funcs, nil
}
