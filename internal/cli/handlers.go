package cli

import (
	"devsyringe/internal/exceptions"
	process "devsyringe/internal/process"
	"fmt"
)

func printProcessList(pm *process.ProcManager) {
	processes := pm.GetProcesses()
	for _, proc := range processes {
		fmt.Printf("%s\t%d\t(%s)\n", proc.Title, proc.PID, proc.Status)
	}
	fmt.Printf("Total: %d prcesses.\n", len(processes))
}

func stopProcessHandler(title string, stopAll bool, pm *process.ProcManager) {
	if stopAll {
		pm.StopAllProcesses()
		fmt.Printf("Stop all processes.\n")
		return
	}

	err := pm.StopProcess(title)
	exceptions.Print(err)
	if err == nil {
		fmt.Printf("Stop process with title %s.\n", title)
	}
}

func deleteProcessHandler(title string, deleteAll bool, pm *process.ProcManager) {
	if deleteAll {
		pm.DeleteAllProcesses()
		fmt.Printf("Delete all processes.\n")
		return
	}
	err := pm.DeleteProcess(title)
	exceptions.Print(err)
	if err == nil {
		fmt.Printf("Delete process with title %s.\n", title)
	}
}

func printProcessLogs(title string, pm *process.ProcManager) {
	logs, err := pm.GetProcessLogs(title)
	exceptions.Print(err)
	fmt.Printf("%s\n", logs)
}
