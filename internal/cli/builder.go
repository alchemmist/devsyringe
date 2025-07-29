package cli

import (
	procmng "devsyringe/internal/proc"

	"github.com/spf13/cobra"
)

func BuildCli(pm *procmng.ProcManager) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "dsy",
		Short: "Developer tool for hard injection to files.",
	}

	rootCmd.AddCommand(injectCmd(pm))
	rootCmd.AddCommand(listCmd(pm))
	rootCmd.AddCommand(stopCmd(pm))
	rootCmd.AddCommand(deleteCmd(pm))

	return rootCmd
}
