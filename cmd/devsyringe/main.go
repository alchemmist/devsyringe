package main

import (
	"errors"
	"flag"
	"log"
	"os"

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

func main() {
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

	log.Print(config.Serums)
}
