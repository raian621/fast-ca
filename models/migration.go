package models

import (
	"database/sql"
	"log"
)

// define database migrations here:
var MIGRATIONS = []MigrationModel{}

// Represents a database schema migration
type MigrationModel struct {
  Id           int
  // Name of the migration (ex. `Add Users Table`). The name must be unique
  // (and ideally descriptive)
  Name         string
  // SQL statement(s) to run for the migration 
  SqlStatement string
}

func NewMigration(name, sqlStatement string) *MigrationModel {
  return &MigrationModel{
    Name: name,
    SqlStatement: sqlStatement,
  }
}

func (m *MigrationModel) Insert(db Queryable) error {
  _, err := db.Exec(
    `INSERT INTO migrations (name) VALUES (?)`,
    m.Name,
  )
  return err
}

func (m *MigrationModel) Apply(db Queryable) error {
  _, err := db.Exec(m.SqlStatement)
  return err
}

// Apply migrations to the database schema. If any migration fails, this 
// function will return the error describing the failure. 
func ApplyMigrations(db *sql.DB, migrations []MigrationModel) error {
  _, err := db.Exec(
    "CREATE TABLE IF NOT EXISTS migrations (" +
    " id   SERIAL PRIMARY KEY," +
    " name VARCHAR(200) NOT NULL" +
    ")",
  )
  if err != nil {
    return err
  }

  rows, err := db.Query("SELECT name FROM migrations")
  if err != nil {
    return err
  }
  applied := make(map[string]bool, 0)

  for rows.Next() {
    var name string
    if err := rows.Scan(&name); err != nil {
      return err
    }
    applied[name] = true
  }
  if err := rows.Err(); err != nil {
    return err
  }

  for _, migration := range migrations {
    if _, ok := applied[migration.Name]; ok {
      // skip migration if it's already been applied
      log.Printf("Migration `%s` already applied\n", migration.Name)
      continue
    }
    
    log.Printf("Applying migration `%s`...\n", migration.Name)
    
    tx, err := db.Begin()
    if err != nil {
      return err
    }
    if err := migration.Apply(tx); err != nil {
      return err
    }
    if err := migration.Insert(tx); err != nil {
      return err
    }
    if err := tx.Commit(); err != nil {
      return err
    }

    log.Printf("Migration `%s` successful!\n", migration.Name)
  }

  return nil
}

