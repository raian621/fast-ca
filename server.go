package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/raian621/fast-ca/api"
)

type Server struct{}

var _ api.ServerInterface = (*Server)(nil)

func (s *Server) GetCertificateCertId(ctx echo.Context, certId int) error {
	return nil
}

func (s *Server) GetCertificateList(ctx echo.Context, params api.GetCertificateListParams) error {
	return nil
}

func (s *Server) PostCertificate(ctx echo.Context) error {
	return nil
}

func (s *Server) GetOpenapiYml(ctx echo.Context) error {
	file, err := os.Open("./openapi.yml")
	defer func() {
    if err := file.Close(); err != nil {
      log.Println(err)
    }
  }()
	if err != nil {
		return err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	ctx.Response().Header().Set("Content-Type", "application/yaml")
	return ctx.String(http.StatusOK, string(data))
}

func (s *Server) GetDocs(ctx echo.Context) error {
	file, err := os.Open("./docs.html")
	defer func() {
    if err := file.Close(); err != nil {
      log.Println(err)
    }
  }()
	if err != nil {
		return err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	return ctx.HTML(http.StatusOK, string(data))
}

func (s *Server) GetCaCaIdFullchain(ctx echo.Context, caId int) error {
	return nil
}

func (s *Server) GetCaList(ctx echo.Context, params api.GetCaListParams) error {
	return nil
}

func (s *Server) PostCa(ctx echo.Context) error {
	return nil
}

func (s *Server) PostSignin(ctx echo.Context) error {
	var data api.UserCredentials
	if err := json.NewDecoder(ctx.Request().Body).Decode(&data); err != nil {
		return err
	}

	return ctx.NoContent(200)
}

func (s *Server) PostSignout(ctx echo.Context) error {
	return nil
}
