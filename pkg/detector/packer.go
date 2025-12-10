package detector

import (
	"github.com/dmitry-lun/Mali/internal/domain/entity"
)

func DetectPackers(sections []entity.Section) []string {
	packers := []string{}
	packerMap := make(map[string]bool)

	for _, sec := range sections {
		switch sec.Name {
		case "UPX0", "UPX1":
			if !packerMap["UPX"] {
				packers = append(packers, "UPX")
				packerMap["UPX"] = true
			}
		case ".themida":
			if !packerMap["Themida"] {
				packers = append(packers, "Themida")
				packerMap["Themida"] = true
			}
		case ".vmp0", ".vmp1":
			if !packerMap["VMProtect"] {
				packers = append(packers, "VMProtect")
				packerMap["VMProtect"] = true
			}
		case ".aspack":
			if !packerMap["ASPack"] {
				packers = append(packers, "ASPack")
				packerMap["ASPack"] = true
			}
		}
	}
	return packers
}
