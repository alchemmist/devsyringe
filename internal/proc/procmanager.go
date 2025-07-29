package procmng

import (
	"devsyringe/internal/exceptions"
	"devsyringe/internal/paths"
	"devsyringe/internal/utils"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"syscall"

	"github.com/jmoiron/sqlx"
)

type ProcessSearchFilter struct{}

type ProcessSearchPredicate func(p *Process) bool

func (f ProcessSearchFilter) ByTitle(title string) ProcessSearchPredicate {
	return func(p *Process) bool {
		return p.Title == title
	}
}

func (f ProcessSearchFilter) ByPID(pid int) ProcessSearchPredicate {
	return func(p *Process) bool {
		return p.PID == pid
	}
}

var filter = ProcessSearchFilter{}

type ProcManager struct {
	mu        sync.RWMutex
	processes []*Process
	db        *sqlx.DB
}

func NewProcManager(db *sqlx.DB) *ProcManager {
	pm := &ProcManager{
		db: db,
	}
	pm.loadProcesses()
	return pm
}

func (pm *ProcManager) loadProcesses() {
	var processes []*Process
	err := pm.db.Select(&processes, `SELECT * FROM processes;`)
	exceptions.Check(err)

	pm.mu.Lock()
	defer pm.mu.Unlock()

	pm.processes = processes
}

func (pm *ProcManager) saveProcessToDB(proc *Process) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	_, err := pm.db.NamedExec(`INSERT OR REPLACE INTO processes 
        (title, pid, log_file, status, command) 
        VALUES (:title, :pid, :log_file, :status, :command);`, proc)
	exceptions.Check(err)
}

func (pm *ProcManager) deleteProcessFromDB(proc *Process) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	_, err := pm.db.Exec(`DELETE FROM processes WHERE title = ? AND pid = ?;`, proc.Title, proc.PID)
	exceptions.Check(err)
}

func (pm *ProcManager) syncProcessesStatus() {
	pm.mu.Lock()
	var toUpdate []*Process
	for _, proc := range pm.processes {
		if !proc.IsAlive() {
			proc.Status = Stoped
			toUpdate = append(toUpdate, proc)
		}
	}
	pm.mu.Unlock()

	for _, proc := range toUpdate {
		pm.saveProcessToDB(proc)
	}
}

func (pm *ProcManager) findProcess(searchPredicate ProcessSearchPredicate) (*Process, error) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	for _, proc := range pm.processes {
		if searchPredicate(proc) {
			return proc, nil
		}
	}
	return nil, fmt.Errorf("no process with this parameter")
}

func (pm *ProcManager) DeleteProcess(title string) {
	proc, err := pm.findProcess(filter.ByTitle(title))
	if err != nil {
		fmt.Printf("No process with title %s.\n", title)
		return
	}

	pm.mu.Lock()
	if proc.IsAlive() {
		proc.Stop()
	}
	pm.processes = utils.Remove(pm.processes, proc)
	pm.mu.Unlock()

	pm.deleteProcessFromDB(proc)
}

func (pm *ProcManager) StopProcess(title string) {
	proc, err := pm.findProcess(filter.ByTitle(title))
	if err != nil {
		fmt.Printf("No process with title %s.\n", title)
		return
	}

	if !proc.IsAlive() {
		fmt.Printf("Process %s already stoped.\n", title)
		return
	}

	pm.mu.Lock()
	defer pm.mu.Unlock()
	proc.Stop()
}

func (pm *ProcManager) GetProcesses() []*Process {
	pm.syncProcessesStatus()

	pm.mu.RLock()
	defer pm.mu.RUnlock()

	return pm.processes
}

func (pm *ProcManager) StartProcess(title string, command string) {
	var proc *Process
	var err error

	if proc, err = pm.findProcess(filter.ByTitle(title)); err == nil {
		proc.Restart()
	} else {
		outputFile := filepath.Join(paths.GetLogsDirectory(),
			fmt.Sprintf("process_%s.log", title))
		logFile, err := os.Create(outputFile)
		exceptions.Check(err)

		defer logFile.Close()

		cmd := exec.Command("sh", "-c", command)
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Setsid: true,
		}

		cmd.Stdout = logFile
		cmd.Stderr = logFile

		err = cmd.Start()
		exceptions.Check(err)

		proc = &Process{
			Title:   title,
			PID:     cmd.Process.Pid,
			LogFile: outputFile,
			Status:  Active,
			Command: command,
		}
	}
	pm.saveProcessToDB(proc)
}
