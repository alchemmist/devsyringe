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

func (pm *ProcManager) loadProcesses() error {
	var processes []*Process
	err := pm.db.Select(&processes, `SELECT * FROM processes;`)
	exceptions.Check(err)

	pm.mu.Lock()
	defer pm.mu.Unlock()

	pm.processes = processes
	return err
}

func (pm *ProcManager) saveProcess(proc Process) error {
	_, err := pm.db.NamedExec(`INSERT OR REPLACE INTO processes 
        (title, pid, log_file, status, command) 
        VALUES (:title, :pid, :log_file, :status, :command);`, &proc)
	return err
}

func (pm *ProcManager) KillProcess(title string) {

}

func (pm *ProcManager) StopProcess(title string) {

}

func (pm *ProcManager) GetProcesses() []*Process {
	err := pm.loadProcesses()
	exceptions.Check(err)

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

	pm.mu.Lock()
	defer pm.mu.Unlock()

	err = pm.saveProcess(process)
	exceptions.Check(err)
}
