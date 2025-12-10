package reports

import (
	"encoding/json"
	"os"

	"github.com/dmitry-lun/Mali/internal/domain/entity"
)

func GenerateJSON(peFile *entity.PEFile, outputPath string) error {
	jsonData, err := json.MarshalIndent(peFile, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(outputPath, jsonData, 0644)
}
