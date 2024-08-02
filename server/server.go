package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/raian621/fast-ca/models"
	_ "modernc.org/sqlite"
)

type Server struct {
	db *sql.DB
	e  *echo.Echo
}

func (s *Server) DbConn() (*sql.Conn, error) {
	return s.db.Conn(context.Background())
}

func NewServer() (*Server, error) {
  db, err := sql.Open("sqlite", "fast-ca.db")
	if err != nil {
		return nil, err
	}
  if err := models.ApplyMigrations(db, models.MIGRATIONS); err != nil {
    return nil, err
  }

	e := echo.New()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogMethod: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Printf("%v %v [%v] -> %v\n", v.Method, v.URI, c.RealIP(), v.Status)
			return nil
		},
	}))
	e.Static("/", "public")

	e.GET("/", func(c echo.Context) error {
		return c.File("views/index.html")
	})

	e.GET("/ca", func(c echo.Context) error {
		return c.String(http.StatusAccepted, "Hello, CA!")
	})

	e.GET("/ca/new", func(c echo.Context) error {
		return c.File("views/ca/new.html")
	})

	e.POST("/ca/new", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "/ca")
	})

	e.GET("/certificate/new", func(c echo.Context) error {
		return c.File("views/certificate/new")
	})

	e.POST("/certificate/new", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "/certificate")
	})

	e.GET("/login", func(c echo.Context) error {
		return c.File("views/login.html")
	})

	e.POST("/login", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "/")
	})

	return &Server{
		db: db,
		e:  e,
	}, err
}

func (s *Server) Listen(addr string, port int) {
	if err := s.e.Start(fmt.Sprintf("%s:%d", addr, port)); err != nil {
		if dbErr := s.db.Close(); dbErr != nil {
			log.Println(err)
		}
		log.Fatalln(err)
	}
}
