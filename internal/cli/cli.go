package cli

import (
	"devsyringe/internal/config"
	procmng "devsyringe/internal/proc"
	"fmt"

	"github.com/spf13/cobra"
)

func BuildCli(pm *procmng.ProcManager) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "dsy",
		Short: "Developer tool for hard injection to files",
	}

	var configPath string
	var injectCmd = &cobra.Command{
		Use:   "inject",
		Short: "Start a injection based on some config",
		Run: func(cmd *cobra.Command, args []string) {
			conf := config.ParseConfig(configPath)
			config.ProcessingConfig(conf, pm)
		},
	}

	injectCmd.Flags().StringVarP(&configPath, "config", "c", "devsyringe.yaml",
		"The config .yaml file for devsyring")

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "Show dynamic list of running processes",
		Run: func(cmd *cobra.Command, args []string) {
			PrintProcessList(pm)
		},
	}

	rootCmd.AddCommand(injectCmd)
	rootCmd.AddCommand(listCmd)

	return rootCmd
}

func PrintProcessList(pm *procmng.ProcManager) {
	processes := pm.GetProcesses()
	for _, proc := range processes {
		fmt.Printf("%s\t%d\t(%s)\n", proc.Title, proc.PID, proc.Status)
	}
}
