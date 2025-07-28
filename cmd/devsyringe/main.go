package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Serums map[string]Serm `yaml:"serums"`
}

type Serm struct {
	Source  string            `yaml:"source"`
	Mask    string            `yaml:"mask"`
	Targets map[string]Target `yaml:"targets"`
}

type Target struct {
	Path  string   `yaml:"path"`
	Clues []string `yaml:"clues"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseConfig() *Config {
	configPath := flag.String("config", "devsyringe.yaml", "Path to .yaml config file")

	flag.Parse()

	if _, err := os.Stat(*configPath); errors.Is(err, os.ErrNotExist) {
		log.Printf("File %s does not exist", *configPath)
		os.Exit(1)
	}

	configFile, err := os.ReadFile(*configPath)
	check(err)

	var config Config

	err = yaml.Unmarshal([]byte(configFile), &config)
	check(err)

	return &config
}

func processingConfig(config *Config) {
	for title, surm := range config.Serums {
		home, err := os.UserHomeDir()
		check(err)

		outputDir := filepath.Join(home, ".local", "share", "devsyringe")
		err = os.MkdirAll(outputDir, 0755)
		check(err)

		outputFile := filepath.Join(outputDir, fmt.Sprintf("process_%s.log", title))
		logFile, err := os.Create(outputFile)
		check(err)
		defer logFile.Close()

		cmd := exec.Command("sh", "-c", surm.Source)

		cmd.SysProcAttr = &syscall.SysProcAttr{
			Setsid: true,
		}

		cmd.Stdout = logFile
		cmd.Stderr = logFile

		err = cmd.Start()
		check(err)

		log.Printf("Процесс %s запущен с PID %d. Лог в process.log\n", title, cmd.Process.Pid)
	}
}

func main() {
	config := parseConfig()
	processingConfig(config)
}
