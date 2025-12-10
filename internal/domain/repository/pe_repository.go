package repository

import "github.com/dmitry-lun/Mali/internal/domain/entity"

type PERepository interface {
	Parse(path string) (*entity.PEFile, error)
}
