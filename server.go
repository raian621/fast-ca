package main

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/raian621/fast-ca/api"
)

type Server struct {}

var _ api.ServerInterface = (*Server)(nil)

func (s *Server) GetCertificateCertId(ctx echo.Context, certId int) error {
  return nil
}

func (s *Server) GetCertificates(ctx echo.Context, params api.GetCertificatesParams) error {
  return nil
}

func (s *Server) PostCertificate(ctx echo.Context) error {
  return nil
}

func (s *Server) GetOpenapiYml(ctx echo.Context) error {
  file, err := os.Open("./openapi.yml")
  defer file.Close()
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

func (s *Server) GetRedoc(ctx echo.Context) error {
  file, err := os.Open("./redoc.html")
  defer file.Close()
  if err != nil {
    return err
  }

  data, err := io.ReadAll(file)
  if err != nil {
    return err
  }

  return ctx.HTML(http.StatusOK, string(data))
}
