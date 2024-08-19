package database

import (
	"os"

	"github.com/raian621/fast-ca/util"
)

// returns the ID of the created user upon success
func CreateUser(username, email, passhash string) (int, error) {
	_, err := db.Exec(
		"insert into users (username, email, passhash) values (?, ?, ?)",
		username,
		email,
		passhash,
	)
	if err != nil {
		return -1, err
	}
	row := db.QueryRow(
		"select id from users where username=? AND email=?",
		username,
		email,
	)
	var userId int
	err = row.Scan(&userId)
	return userId, err
}

func BootstrapAdminUser() error {
	username := "admin"
	password := "password"
	email := "admin@localhost"

	if envUsername, present := os.LookupEnv("ADMIN_USERNAME"); present {
		username = envUsername
	}
	if envPassword, present := os.LookupEnv("ADMIN_PASSWORD"); present {
		password = envPassword
	}
	if envEmail, present := os.LookupEnv("ADMIN_EMAIL"); present {
		email = envEmail
	}

	passhash, err := util.HashPassword(password)
	if err != nil {
		return nil
	}
	_, err = CreateUser(username, email, passhash)
	return err
}
