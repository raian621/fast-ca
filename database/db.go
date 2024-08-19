package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/raian621/fast-ca/util"
)

var (
	db *sql.DB

	dbPath = "./sqlite.db"
)

func NewDB() (err error) {
	absDbPath, err := util.RelativeToAbsolutePath(dbPath)
	if err != nil {
		return err
	}
	log.Printf("Using database file at `%s`\n", absDbPath)

	dbExists := true
	if _, err := os.Stat(absDbPath); os.IsNotExist(err) {
		dbExists = false
	}

  db, err = sql.Open("sqlite3", sqliteConnString(absDbPath))
  if err != nil {
    return err
  }
  if !dbExists {
		if err := createMigrationsTable(); err != nil {
			return err
		}
	}

	if err := applyMigrations(); err != nil {
		return err
	}

	if !dbExists {
		log.Println("Initializing admin user...")
		return BootstrapAdminUser()
	}

	return nil
}

func sqliteConnString(dbPath string) string {
	return fmt.Sprintf("file:%s", dbPath)
}
