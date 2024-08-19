package database

import "log"

type migration struct {
	name string
	sql  string
}

var migrations = []migration{
	{
		name: "add users table",
		sql: ("create table users (\n" +
			"  id       INTEGER      PRIMARY KEY AUTOINCREMENT,\n" +
			"  username VARCHAR(256) NOT NULL UNIQUE,\n" +
			"  email    VARCHAR(256) NOT NULL UNIQUE,\n" +
			"  passhash CHAR(97)     NOT NULL UNIQUE\n" +
			")\n"),
	},
}

func createMigrationsTable() error {
	_, err := db.Exec(
		"create table migrations (\n" +
			"  id   INTEGER      PRIMARY KEY AUTOINCREMENT,\n" +
			"  name VARCHAR(200) NOT NULL UNIQUE\n" +
			")",
	)
	return err
}

func applyMigrations() error {
	for _, migration := range migrations {
		var count int
		row := db.QueryRow("select count(*) from migrations where name=?", migration.name)
		if err := row.Scan(&count); err != nil {
			return err
		}
		if count == 1 {
			log.Printf("Skipping migration `%s`...\n", migration.name)
			continue
		}

		log.Printf("Applying migration `%s`...\n", migration.name)
		tx, err := db.Begin()
		if err != nil {
			return err
		}
		if _, err := tx.Exec(migration.sql); err != nil {
			return err
		}
		if _, err := tx.Exec(
			"insert into migrations (name) values (?)",
			migration.name,
		); err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
	}

	return nil
}
