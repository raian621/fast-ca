package database

import (
	"log"
	"os"
	"testing"
)

var testDbPath = "./sqlite.test.db"

func TestMain(m *testing.M) {
  setup()
  code := m.Run()
  teardown()
  os.Exit(code)
}

func setup() {
  if _, err := NewDB(testDbPath); err != nil { 
    log.Fatalf("Unexpected error occurred during test setup: %v\n", err)
  }
}

func teardown() {
  if err := os.Remove(testDbPath); err != nil { 
    log.Fatalf("Unexpected error occurred during test teardown: %v\n", err)
  }
}

func assertEqual[T comparable](t *testing.T, want, got T) {
  if want != got {
    t.Errorf("expected `%v`, got `%v`", want, got)
  }
}

func TestBootstrapAdminUser(t *testing.T) {
  if err := BootstrapAdminUser(); err != nil {
    t.Fatal(err)
  }

  row := db.QueryRow("select id, username, email, passhash from users where username='admin'")
  var (
    id       int
    username string
    email    string
    passhash string
  )
  if err := row.Scan(&id, &username, &email, &passhash); err != nil {
    t.Fatal(err)
  }

  assertEqual(t, 1, id)
  assertEqual(t, "admin", username)
  assertEqual(t, "admin@localhost", email)
}
