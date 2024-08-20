package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

// Creates a new SQLite database. If the database is being initialized for the
// first time, it's migrations are applied and it will return true to indicate
// that the database needs to be further bootstrapped (needs an Admin user,
// default CA, etc.)
func NewDB(dbPath string) (needsBootstrapping bool, err error) {
	log.Printf("Using database file at `%s`\n", dbPath)

	dbExists := true
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		dbExists = false
	}

  db, err = sql.Open("sqlite3", sqliteConnString(dbPath))
  if err != nil {
    return false, err
  }
  if !dbExists {
		if err := createMigrationsTable(); err != nil {
			return false, err
		}
	}

	if err := applyMigrations(); err != nil {
		return false, err
	}

	return !dbExists, nil
}

func sqliteConnString(dbPath string) string {
	return fmt.Sprintf("file:%s", dbPath)
}
