package main

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=oapi-codegen.yml openapi.yml

import "fmt"

func main() {
  fmt.Println("Hello, world!")
}
