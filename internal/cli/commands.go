package cli

import (
	"devsyringe/internal/config"
	procmng "devsyringe/internal/proc"

	"github.com/spf13/cobra"
)

func injectCmd(pm *procmng.ProcManager) *cobra.Command {
	var configPath string
	var injectCmd = &cobra.Command{
		Use:   "inject",
		Short: "Start a injection based on some config.",
		Run: func(cmd *cobra.Command, args []string) {
			conf := config.ParseConfig(configPath)
			config.ProcessingConfig(conf, pm)
		},
	}
	injectCmd.Flags().StringVarP(&configPath, "config", "c", "devsyringe.yaml",
		"The config .yaml file for devsyring")
	return injectCmd
}

func listCmd(pm *procmng.ProcManager) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "Show dynamic list of running processes.",
		Run: func(cmd *cobra.Command, args []string) {
			PrintProcessList(pm)
		},
	}
	return listCmd
}

func stopCmd(pm *procmng.ProcManager) *cobra.Command {
	stopCmd := &cobra.Command{
		Use:   "stop [title]",
		Short: "Stop process with [title], but save logs and save in list.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			StopProcess(args[0], pm)
		},
	}
	return stopCmd
}

func deleteCmd(pm *procmng.ProcManager) *cobra.Command {
	stopCmd := &cobra.Command{
		Use: "delete [title]",
		Short: "If not stoped, stop. Then, delete process " +
			"with [title] from list and delete all logs.",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			DeleteProcess(args[0], pm)
		},
	}
	return stopCmd
}
