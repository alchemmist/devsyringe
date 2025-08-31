package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/alchemmist/devsyringe/internal/cli"
	"github.com/alchemmist/devsyringe/internal/paths"
	process "github.com/alchemmist/devsyringe/internal/process"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func LoadSchema(db *sqlx.DB, path string) error {
	schemaBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(schemaBytes))
	return err
}

func InitDB(path string) *sqlx.DB {
	db, err := sqlx.Open("sqlite3", path)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func main() {
	database := InitDB(filepath.Join(paths.GetDataDirectory(), "data.sqlite"))
	LoadSchema(database, "db/schema.sql")

	processManager := process.NewProcManager(database)
	app := cli.BuildCli(processManager)
	app.Execute()
}
