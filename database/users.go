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

func DeleteUser(userId int) error {
  _, err := db.Exec("delete from users where id=?", userId)
  return err
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
  
  salt, err := util.GenerateSalt()
	if err != nil {
		return nil
	}
	passhash := util.HashPassword(password, salt)
	_, err = CreateUser(username, email, passhash)
	return err
}

func ValidateCredentials(usernameOrEmail, password string) (bool, error) {
  row := db.QueryRow(
    "select passhash from users where username=? or email=?",
    usernameOrEmail,
    usernameOrEmail,
  )
  var passhash string
  if err := row.Scan(&passhash); err != nil {
    return false, err
  }

  return util.ValidatePassword(password, passhash)
}
