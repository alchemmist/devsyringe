package process

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/alchemmist/devsyringe/internal/exceptions"
	"github.com/alchemmist/devsyringe/internal/paths"
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

func (p *Process) Stop() {
	proc, err := os.FindProcess(p.PID)
	exceptions.Check(err)

	err = proc.Signal(syscall.SIGTERM)
	exceptions.Check(err)

	p.Status = Stoped
}

func (p *Process) Restart() {
	if p.IsAlive() {
		p.Stop()
	}

	outputFile := filepath.Join(paths.GetLogsDirectory(),
		fmt.Sprintf("process_%s.log", p.Title))
	logFile, err := os.Create(outputFile)
	exceptions.Check(err)

	cmd := exec.Command("sh", "-c", p.Command)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}

	cmd.Stdout = logFile
	cmd.Stderr = logFile

	err = cmd.Start()
	exceptions.Check(err)

	p.PID = cmd.Process.Pid
	p.Status = Active
}

func (p *Process) GetLogs() string {
	data, err := os.ReadFile(p.LogFile)
	exceptions.Check(err)

	logs := string(data)
	return logs
}
