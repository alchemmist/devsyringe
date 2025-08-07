package cli

import (
	"devsyringe/internal/cli/tui"
	process "devsyringe/internal/process"

	"github.com/spf13/cobra"
)

const syringeArt = `
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⠝⡄⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠤⡀⠀⠀⠀⠀⣘⡴⡀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢈⡄⠡⣀⣤⣶⣿⣿⣷⡱⡀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⣴⠚⢉⠈⡄⢹⣿⣿⠿⠛⠉⠑⡡⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⢴⢺⡿⠋⢭⠀⡘⡄⠘⡀⢫⠀⠀⠀⠀⠀⠑⠃
⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⠁⠀⡀⠁⢣⠐⡈⢆⡱⠜⠊⠑⣀⡆⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⣠⢊⣇⠀⠱⣘⡤⠗⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⣀⠠⠐⠉⠀⠁⠈⠓⠊⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠈⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`

func BuildCli(pm *process.ProcManager) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "dsy",
		Short: "\nDeveloper tool for hard injection to files.\n" + syringeArt,
		Run: func(cmd *cobra.Command, args []string) {
			tui.Tui(pm)
		},
	}

	rootCmd.AddCommand(injectCmd(pm))
	rootCmd.AddCommand(listCmd(pm))
	rootCmd.AddCommand(stopCmd(pm))
	rootCmd.AddCommand(deleteCmd(pm))
	rootCmd.AddCommand(logsCmd(pm))

	return rootCmd
}
