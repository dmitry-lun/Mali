package main

import (
	"github.com/dmitry-lun/Mali/internal/domain/usecase"
	"github.com/dmitry-lun/Mali/internal/infrastructure/repository"
	"github.com/dmitry-lun/Mali/internal/presentation/cli"
	"github.com/dmitry-lun/Mali/pkg/file"
)

func main() {
	fileReader := file.NewFileReader()
	peRepo := repository.NewPERepository(fileReader)
	analyzeUseCase := usecase.NewAnalyzePEUseCase(peRepo)

	analysisHandler := cli.NewAnalysisHandler(analyzeUseCase)

	rootCmd := cli.NewRootCommand()
	rootCmd.AddCommand(analysisHandler.CreateCommand())
	rootCmd.Execute()
}
