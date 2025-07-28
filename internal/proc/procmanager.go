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

type ProcessInfo struct {
	Title   string     `db:"title"`
	PID     int        `db:"pid"`
	LogFile string     `db:"log_file"`
	Status  ProcStatus `db:"status"`
	Command string     `db:"command"`
}

type ProcManager struct {
	mu        sync.Mutex
	processes map[string]*ProcessInfo
	db        *sqlx.DB
}

func NewProcManager(db *sqlx.DB) *ProcManager {
	return &ProcManager{
		processes: make(map[string]*ProcessInfo),
		db:        db,
	}
}

func (pm *ProcManager) saveProcInfo(procInfo ProcessInfo) error {
	_, err := pm.db.NamedExec(`INSERT OR REPLACE INTO processes 
        (title, pid, log_file, status, command) 
        VALUES (:title, :pid, :log_file, :status, :command)`, &procInfo)
	return err
}

func (pm *ProcManager) KillProcess(title string) {

}

func (pm *ProcManager) StopProcess(title string) {

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

	processInfo := ProcessInfo{
		Title:   title,
		PID:     cmd.Process.Pid,
		LogFile: outputFile,
		Status:  Active,
		Command: command,
	}

	err = pm.saveProcInfo(processInfo)
	exceptions.Check(err)
}
