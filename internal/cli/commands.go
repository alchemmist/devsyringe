package cli

import (
	"fmt"

	"github.com/alchemmist/devsyringe/internal/config"
	process "github.com/alchemmist/devsyringe/internal/process"

	"github.com/spf13/cobra"
)

func injectCmd(pm *process.ProcManager) *cobra.Command {
	var configPath string
	var verboseLogs bool = false
	var injectCmd = &cobra.Command{
		Use:   "inject",
		Short: "Start a injection based on some config.",
		Run: func(cmd *cobra.Command, args []string) {
			conf := config.ParseConfig(configPath)
			config.ProcessingConfig(conf, verboseLogs, pm)
		},
	}
	injectCmd.Flags().StringVarP(&configPath, "config", "c", "devsyringe.yaml",
		"The config .yaml file for devsyring")
	injectCmd.Flags().BoolVarP(&verboseLogs, "verbose", "v", false,
		"Show source and replaced line in every target")
	return injectCmd
}

func listCmd(pm *process.ProcManager) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "Show dynamic list of running processes.",
		Run: func(cmd *cobra.Command, args []string) {
			printProcessList(pm)
		},
	}
	return listCmd
}

func stopCmd(pm *process.ProcManager) *cobra.Command {
	var stopAll bool = false
	stopCmd := &cobra.Command{
		Use:   "stop [title]",
		Short: "Stop process with [title], but save logs and save in list.",
		Args: func(cmd *cobra.Command, args []string) error {
			stopAll, _ := cmd.Flags().GetBool("all")
			if !stopAll && len(args) < 1 {
				return fmt.Errorf("You must provide a [title] unless --all is specified.")
			}
			if stopAll && len(args) > 0 {
				return fmt.Errorf("Cannot provide [title] when --all is used.")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				stopProcessHandler("", stopAll, pm)
			} else {
				stopProcessHandler(args[0], stopAll, pm)
			}
		},
	}
	stopCmd.Flags().BoolVarP(&stopAll, "all", "", false,
		"Stop all process from list.")
	return stopCmd
}

func deleteCmd(pm *process.ProcManager) *cobra.Command {
	var deleteAll bool = false
	deleteCmd := &cobra.Command{
		Use: "delete [title]",
		Short: "If not stoped, stop. Then, delete process " +
			"with [title] from list and delete all logs.",
		Args: func(cmd *cobra.Command, args []string) error {
			deleteAll, _ := cmd.Flags().GetBool("all")
			if !deleteAll && len(args) < 1 {
				return fmt.Errorf("You must provide a [title] unless --all is specified.")
			}
			if deleteAll && len(args) > 0 {
				return fmt.Errorf("Cannot provide [title] when --all is used.")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				deleteProcessHandler("", deleteAll, pm)
			} else {
				deleteProcessHandler(args[0], deleteAll, pm)
			}
		},
	}
	deleteCmd.Flags().BoolVarP(&deleteAll, "all", "", false,
		"Delete (and stop) all process from list.")
	return deleteCmd
}

func logsCmd(pm *process.ProcManager) *cobra.Command {
	logsCmd := &cobra.Command{
		Use:   "logs [title]",
		Short: "Show logs from process with [title].",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			printProcessLogs(args[0], pm)
		},
	}
	return logsCmd
}
