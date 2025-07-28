package procmng

import (
	"devsyringe/internal/exceptions"
	"devsyringe/internal/paths"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"syscall"

	"github.com/jmoiron/sqlx"
)

type ProcStatus int

const (
	Stoped ProcStatus = iota
	Active
)

var statusName = map[ProcStatus]string{
	Stoped: "stoped",
	Active: "active",
}

func (ps ProcStatus) String() string {
	return statusName[ps]
}

type Process struct {
	Title   string     `db:"title"`
	PID     int        `db:"pid"`
	LogFile string     `db:"log_file"`
	Status  ProcStatus `db:"status"`
	Command string     `db:"command"`
}

func (p *Process) IsAlive() bool {
	proc, err := os.FindProcess(p.PID)
	if err != nil {
		return false
	}
	err = proc.Signal(syscall.Signal(0))
	return err == nil
}

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

func (pm *ProcManager) saveProcess(proc Process) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	_, err := pm.db.NamedExec(`INSERT OR REPLACE INTO processes 
        (title, pid, log_file, status, command) 
        VALUES (:title, :pid, :log_file, :status, :command);`, &proc)
	exceptions.Check(err)
}

func (pm *ProcManager) syncProcessesStatus() {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	for _, proc := range pm.processes {
		if !proc.IsAlive() {
			proc.Status = Stoped
		}
	}
}

func (pm *ProcManager) DeleteProcess(title string) {

}

func (pm *ProcManager) StopProcess(title string) {

}

func (pm *ProcManager) GetProcesses() []*Process {
	pm.syncProcessesStatus()

	pm.mu.RLock()
	defer pm.mu.RUnlock()

	return pm.processes
}

func (pm *ProcManager) NewProcess(title string, command string) {
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

	process := Process{
		Title:   title,
		PID:     cmd.Process.Pid,
		LogFile: outputFile,
		Status:  Active,
		Command: command,
	}

	pm.saveProcess(process)
	exceptions.Check(err)
}
