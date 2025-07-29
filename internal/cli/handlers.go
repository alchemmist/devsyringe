package cli

import (
	procmng "devsyringe/internal/proc"
	"fmt"
)

func PrintProcessList(pm *procmng.ProcManager) {
	processes := pm.GetProcesses()
	for _, proc := range processes {
		fmt.Printf("%s\t%d\t(%s)\n", proc.Title, proc.PID, proc.Status)
	}
	fmt.Printf("Total: %d prcesses.\n", len(processes))
}

func StopProcess(title string, pm *procmng.ProcManager) {
	pm.StopProcess(title)
}

func DeleteProcess(title string, pm *procmng.ProcManager) {
	pm.DeleteProcess(title)
}
