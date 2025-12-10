package cli

import (
	"fmt"
	"os"

	"github.com/dmitry-lun/Mali/internal/domain/usecase"
	"github.com/dmitry-lun/Mali/pkg/reports"
	"github.com/spf13/cobra"
)

type AnalysisHandler struct {
	analyzeUseCase *usecase.AnalyzePEUseCase
}

func NewAnalysisHandler(analyzeUseCase *usecase.AnalyzePEUseCase) *AnalysisHandler {
	return &AnalysisHandler{
		analyzeUseCase: analyzeUseCase,
	}
}

func (h *AnalysisHandler) CreateCommand() *cobra.Command {
	var filePath string
	var mode string
	var output string

	cmd := &cobra.Command{
		Use:   "analyze",
		Short: "Analyze a PE file",
		Run: func(cmd *cobra.Command, args []string) {
			info, err := h.analyzeUseCase.Execute(filePath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("\n=== PE File Analysis ===")
			fmt.Printf("File: %s\n", info.Path)
			fmt.Printf("Machine: %s\n", info.Machine)
			fmt.Printf("Entry Point: 0x%x\n", info.EntryPoint)
			fmt.Printf("Size: %d bytes\n", info.Size)
			fmt.Printf("Risk Level: %s\n", info.RiskLevel)
			fmt.Printf("MD5: %s\n", info.MD5)
			fmt.Printf("SHA256: %s\n", info.SHA256)

			if len(info.Packers) > 0 {
				fmt.Printf("Detected Packers: %v\n", info.Packers)
			}

			if output != "" {
				if err := reports.GenerateJSON(info, output); err != nil {
					fmt.Fprintf(os.Stderr, "Error generating report: %v\n", err)
					os.Exit(1)
				}
				fmt.Printf("\nReport saved to: %s\n", output)
			}
		},
	}

	cmd.Flags().StringVarP(&filePath, "file", "f", "", "PE file to analyze")
	cmd.Flags().StringVarP(&mode, "mode", "m", "all", "Analysis mode: static|all")
	cmd.Flags().StringVarP(&output, "output", "o", "out.json", "Output file")
	cmd.MarkFlagRequired("file")

	return cmd
}
