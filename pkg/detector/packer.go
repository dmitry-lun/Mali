package detector

import "github.com/dmitry-lun/Mali/internal/domain/entity"

func DetectPackers(sections []entity.Section) []string {
	packers := []string{}
	for _, sec := range sections {
		switch sec.Name {
		case "UPX0", "UPX1":
			packers = append(packers, "UPX")
		case ".themida":
			packers = append(packers, "Themida")
		case ".vmp0":
			packers = append(packers, "VMProtect")
		}
	}
	return packers
}
