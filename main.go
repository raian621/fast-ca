package main

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=oapi-codegen.yml openapi.yml

import (
	"log"
	"os"
	"path"

	"github.com/labstack/echo/v4"
	"github.com/raian621/fast-ca/api"
)

var clientDirectory = "client/dist"

func main() {
  server := &Server{}
  e := echo.New()

  execPath, err := os.Executable()
  if err != nil {
    log.Fatal(err)
  }

  baseDir := path.Dir(execPath)
  absClientDirectory := path.Join(baseDir, clientDirectory)
  if value, present := os.LookupEnv("CLIENT_DIRECTORY"); present {
    absClientDirectory = value
  }

  log.Println(absClientDirectory)
  e.Static("/", absClientDirectory)

  group := e.Group("/api/v1")
  api.RegisterHandlers(group, server)
  log.Fatal(e.Start("0.0.0.0:8080"))
}
