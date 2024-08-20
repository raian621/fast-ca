package main

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=oapi-codegen.yml openapi.yml

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/raian621/fast-ca/database"
	"github.com/raian621/fast-ca/util"
)

var (
	clientDir = "client/dist"
  dbPath = "./sqlite.db"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

  var absDbPath string
  if absDbPath, err := util.RelativeToAbsolutePath(absDbPath); err != nil {
    log.Fatal(absDbPath)
  }
	if needsBootstrapping, err := database.NewDB(absDbPath); err != nil {
		log.Fatal(err)
	} else if needsBootstrapping {
		log.Println("Initializing admin user...")
    if err := database.BootstrapAdminUser(); err != nil {
      log.Fatal(err)
    }
  }

	absClientDirectory, err := util.RelativeToAbsolutePath(clientDir)
	if err != nil {
		log.Fatal(err)
	}

  server := NewServer(absClientDirectory)
	log.Fatal(server.Start("0.0.0.0:8080"))
}

