package main

import (
	"log"
	"os"
	"testing"

	"github.com/raian621/fast-ca/database"
)

var testDbPath = "./sqlite.test.db"

func TestMain(m *testing.M) {
  setup()
  code := m.Run()
  teardown()
  os.Exit(code)
}

func setup() {
  if _, err := database.NewDB(testDbPath); err != nil { 
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

