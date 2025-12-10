package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

type RootCommand struct {
	rootCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCmd := &cobra.Command{
		Use:   "mali",
		Short: "Mali — PE analysis CLI",
		Long:  "CLI for PE File Analysis",
	}

	return &RootCommand{
		rootCmd: rootCmd,
	}
}

func (r *RootCommand) AddCommand(cmd *cobra.Command) {
	r.rootCmd.AddCommand(cmd)
}

func (r *RootCommand) Execute() {
	if err := r.rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
	}
}
