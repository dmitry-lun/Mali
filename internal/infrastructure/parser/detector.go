package parser

import (
	"github.com/dmitry-lun/Mali/internal/domain/entity"
	"github.com/dmitry-lun/Mali/pkg/detector"
)

func DetectPackers(sections []entity.Section) []string {
	return detector.DetectPackers(sections)
}
