package cli

import (
	"devsyringe/internal/exceptions"
	process "devsyringe/internal/process"
	"fmt"
)

func PrintProcessList(pm *process.ProcManager) {
	processes := pm.GetProcesses()
	for _, proc := range processes {
		fmt.Printf("%s\t%d\t(%s)\n", proc.Title, proc.PID, proc.Status)
	}
	fmt.Printf("Total: %d prcesses.\n", len(processes))
}

func StopProcess(title string, pm *process.ProcManager) {
	err := pm.StopProcess(title)
	exceptions.Print(err)
}

func DeleteProcess(title string, pm *process.ProcManager) {
	err := pm.DeleteProcess(title)
	exceptions.Print(err)
}

func PrintProcessLogs(title string, pm *process.ProcManager) {
	logs, err := pm.GetProcessLogs(title)
	exceptions.Print(err)
	fmt.Printf("%s\n", logs)
}
