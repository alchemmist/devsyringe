package cli

import (
	process "devsyringe/internal/process"

	"github.com/spf13/cobra"
)

func BuildCli(pm *process.ProcManager) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "dsy",
		Short: "Developer tool for hard injection to files.",
		Run: func(cmd *cobra.Command, args []string) {
			tui(pm)
		},
	}

	rootCmd.AddCommand(injectCmd(pm))
	rootCmd.AddCommand(listCmd(pm))
	rootCmd.AddCommand(stopCmd(pm))
	rootCmd.AddCommand(deleteCmd(pm))
	rootCmd.AddCommand(logsCmd(pm))

	return rootCmd
}
