package cli

import (
	"os"

	process "github.com/alchemmist/devsyringe/internal/process"
	"github.com/alchemmist/devsyringe/internal/version"

	"github.com/alchemmist/devsyringe/internal/cli/tui"

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

	var showVersion bool

	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Show version")

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if showVersion {
			println("Devsyringe", version.Version)
			os.Exit(0)
		}
	}

	rootCmd.AddCommand(injectCmd(pm))
	rootCmd.AddCommand(listCmd(pm))
	rootCmd.AddCommand(stopCmd(pm))
	rootCmd.AddCommand(deleteCmd(pm))
	rootCmd.AddCommand(logsCmd(pm))

	return rootCmd
}
