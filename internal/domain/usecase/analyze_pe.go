package usecase

import (
	"github.com/dmitry-lun/Mali/internal/domain/entity"
	"github.com/dmitry-lun/Mali/internal/domain/repository"
)

type AnalyzePEUseCase struct {
	peRepo repository.PERepository
}

func NewAnalyzePEUseCase(peRepo repository.PERepository) *AnalyzePEUseCase {
	return &AnalyzePEUseCase{
		peRepo: peRepo,
	}
}

func (uc *AnalyzePEUseCase) Execute(path string) (*entity.PEFile, error) {
	return uc.peRepo.Parse(path)
}
