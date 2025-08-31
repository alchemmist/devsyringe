package paths

import (
	"os"
	"path/filepath"

	"github.com/alchemmist/devsyringe/internal/exceptions"
)

func GetDataDirectory() string {
	home, err := os.UserHomeDir()
	exceptions.Check(err)

	outputDir := filepath.Join(home, ".local", "share", "devsyringe")
	err = os.MkdirAll(outputDir, 0755)
	exceptions.Check(err)

	return outputDir
}

func GetLogsDirectory() string {
	outputDir := filepath.Join(GetDataDirectory(), "logs")
	err := os.MkdirAll(outputDir, 0755)
	exceptions.Check(err)

	return outputDir
}
