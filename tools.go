//go:build tools
// +build tools

package main

import (
	_ "github.com/mitranim/gow"
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
  _ "github.com/golangci/golangci-lint/cmd/golangci-lint"
)
