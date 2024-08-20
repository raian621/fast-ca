package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/raian621/fast-ca/api"
	"github.com/raian621/fast-ca/auth"
)

type Server struct{
  e *echo.Echo
}

var _ api.ServerInterface = (*Server)(nil)

// Create a new server object. It will serve requests to the REST API on the
// `/api/v1` path and will serve client web pages and assets from the root path
// `/`
func NewServer(clientDir string) *Server {
	server := &Server{}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:8080",
			// Swagger OpenAPI viewer
			"https://editor-next.swagger.io/",
			"https://editor.swagger.io/",
		},
		AllowMethods: []string{"GET", "PUT", "POST", "DELETE", "OPTION"},
	}))

	e.Static("/", clientDir)
	group := e.Group("/api/v1")
	api.RegisterHandlers(group, server)
  server.e = e

  return server
}

func (s *Server) Start(addr string) error {
  return s.e.Start(addr)
}

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
	var userCreds api.UserCredentials
	if err := json.NewDecoder(ctx.Request().Body).Decode(&userCreds); err != nil {
    return err
	}

  log.Println(userCreds)
  valid, err := auth.AuthenticateUser(userCreds.Username, userCreds.Password)
  if err != nil {
    return err
  }
  if !valid {
    return ctx.JSON(http.StatusBadRequest, "username or password is incorrect")
  }

	return ctx.NoContent(200)
}

func (s *Server) PostSignout(ctx echo.Context) error {
	return nil
}
