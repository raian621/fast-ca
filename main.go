package main

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=oapi-codegen.yml openapi.yml

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/raian621/fast-ca/api"
	"github.com/raian621/fast-ca/database"
	"github.com/raian621/fast-ca/util"
)

var (
	clientDir = "client/dist"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := database.NewDB(); err != nil {
		log.Fatal(err)
	}

	server := &Server{}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:8080",
			// Swagger OpenAPI editor
			"https://editor-next.swagger.io/",
			"https://editor.swagger.io/",
		},
		AllowMethods: []string{"GET", "PUT", "POST", "DELETE", "OPTION"},
	}))

	absClientDirectory, err := util.RelativeToAbsolutePath(clientDir)
	if err != nil {
		log.Fatal(err)
	}

	e.Static("/", absClientDirectory)

	group := e.Group("/api/v1")
	api.RegisterHandlers(group, server)
	log.Fatal(e.Start("0.0.0.0:8080"))
}
