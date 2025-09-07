package main

import (
	_ "embed"
	"log"
	"path/filepath"

	"github.com/alchemmist/devsyringe/internal/cli"
	"github.com/alchemmist/devsyringe/internal/paths"
	process "github.com/alchemmist/devsyringe/internal/process"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var schemaSQL string

func LoadSchema(db *sqlx.DB, schema string) error {
	_, err := db.Exec(schema)
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
	if err := LoadSchema(database, schemaSQL); err != nil {
		log.Fatalf("failed to load schema: %v", err)
	}

	processManager := process.NewProcManager(database)
	app := cli.BuildCli(processManager)
	app.Execute()
}
