package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/raian621/fast-ca/server"
)

var port = 1234

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("error loading .env file:", err)
	}

	srv, err := server.NewServer()
	if err != nil {
		log.Fatalln(err)
	}
  srv.Listen("", port)
}
