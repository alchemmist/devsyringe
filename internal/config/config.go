package config

import (
	"devsyringe/internal/exceptions"
	"errors"
	"log"
	"os"

	procmng "devsyringe/internal/proc"

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

func ParseConfig(configPath string) *Config {
	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		log.Printf("File %s does not exist", configPath)
		os.Exit(1)
	}

	configFile, err := os.ReadFile(configPath)
	exceptions.Check(err)

	var config Config

	err = yaml.Unmarshal([]byte(configFile), &config)
	exceptions.Check(err)

	return &config
}

func ProcessingConfig(config *Config, pm *procmng.ProcManager) {
	for title, surm := range config.Serums {
		pm.StartProcess(title, surm.Source)
	}
}
