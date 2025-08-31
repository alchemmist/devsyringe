package config

import (
	"bufio"
	"context"
	"devsyringe/internal/exceptions"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	process "devsyringe/internal/process"

	"github.com/goccy/go-yaml"
)

const DefaultSurmMaxTimeout = 3000 * time.Millisecond

type Config struct {
	Serums map[string]*Serm `yaml:"serums"`
}

type Serm struct {
	Source     string            `yaml:"source"`
	Mask       string            `yaml:"mask"`
	Targets    map[string]Target `yaml:"targets"`
	MaxTimeout time.Duration     `yaml:"max-timeout"`
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

	for _, surm := range config.Serums {
		if surm.MaxTimeout == 0 {
			surm.MaxTimeout = DefaultSurmMaxTimeout
		} else {
			surm.MaxTimeout = surm.MaxTimeout * time.Millisecond
		}
	}

	return &config
}

func ProcessingConfig(config *Config, verboseLogs bool, pm *process.ProcManager) {
	for title, surm := range config.Serums {
		proc := pm.StartProcess(title, surm.Source)
		re := regexp.MustCompile(surm.Mask)

		ctx, cancel := context.WithTimeout(context.Background(), surm.MaxTimeout)
		defer cancel()

		logChan := make(chan string, 1)

		go func(ctx context.Context, ch chan<- string) {
			for {
				logs := proc.GetLogs()
				select {
				case <-ctx.Done():
					return
				default:
					if re.MatchString(logs) {
						ch <- logs
						return
					}
				}
				time.Sleep(100 * time.Millisecond)
			}
		}(ctx, logChan)

		select {
		case <-ctx.Done():
			fmt.Printf("The timeout (%d ms) for waiting value in logs for %s has expired. Skip this serm.\n", surm.MaxTimeout, title)
			pm.DeleteProcess(title)
			continue
		case logs := <-logChan:
			value := re.FindSubmatch([]byte(logs))[0]
			if value == nil {
				fmt.Printf("Failed to get value from ouput of %s. Skip this serm.\n", title)
				continue
			}
			for _, target := range surm.Targets {
				file, err := os.Open(target.Path)
				if err != nil {
					fmt.Printf("Error opening file %s: %v.\n", target.Path, err)
					continue
				}
				defer file.Close()

				scanner := bufio.NewScanner(file)
				newFileLines := []string{}
				lineNumber := 0

				for scanner.Scan() {
					lineNumber++
					line := scanner.Text()
					allCluesContained := true
					for _, clue := range target.Clues {
						if !strings.Contains(line, clue) {
							allCluesContained = false
							break
						}
					}
					if allCluesContained {
						updatedLine := re.ReplaceAllString(line, string(value))
						newFileLines = append(newFileLines, updatedLine)
						if verboseLogs {
							fmt.Printf("%s:%d: replace %s to %s\n", target.Path, lineNumber, line, updatedLine)
						} else {
							fmt.Printf("→ Update %s.\n", target.Path)
						}
					} else {
						newFileLines = append(newFileLines, line)
					}
				}
				if err := scanner.Err(); err != nil {
					fmt.Printf("Error while reading to file %s: %v.\n", target.Path, err)
				}
				err = os.WriteFile(target.Path, []byte(strings.Join(newFileLines, "\n")), 0644)
				if err != nil {
					fmt.Printf("Error while writing to file %s: %v.\n", target.Path, err)
					continue
				}
			}
		}
	}
}
