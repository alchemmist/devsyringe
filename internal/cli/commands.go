package cli

import (
	"devsyringe/internal/config"
	process "devsyringe/internal/process"

	"github.com/spf13/cobra"
)

func injectCmd(pm *process.ProcManager) *cobra.Command {
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

func listCmd(pm *process.ProcManager) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "Show dynamic list of running processes.",
		Run: func(cmd *cobra.Command, args []string) {
			PrintProcessList(pm)
		},
	}
	return listCmd
}

func stopCmd(pm *process.ProcManager) *cobra.Command {
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

func deleteCmd(pm *process.ProcManager) *cobra.Command {
	deleteCmd := &cobra.Command{
		Use: "delete [title]",
		Short: "If not stoped, stop. Then, delete process " +
			"with [title] from list and delete all logs.",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			DeleteProcess(args[0], pm)
		},
	}
	return deleteCmd
}

func logsCmd(pm *process.ProcManager) *cobra.Command {
	logsCmd := &cobra.Command{
		Use: "logs [title]",
		Short: "Show logs from process with [title].",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			PrintProcessLogs(args[0], pm)
		},
	}
	return logsCmd
}
