package paths

import (
	"devsyringe/internal/exceptions"
	"os"
	"path/filepath"
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
