package sessions

import (
	"os"

	gsessions "github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type CookieOptions struct {
	Name   string
	Secret string
	MaxAge int
}

const sessionCookieName = "fastca-session"

var store = gsessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func New(c echo.Context) (*gsessions.Session, error) {
	return store.New(c.Request(), sessionCookieName)
}

func Get(c echo.Context) (*gsessions.Session, error) {
	return store.Get(c.Request(), sessionCookieName)
}

func Delete(c echo.Context) error {
	session, err := store.Get(c.Request(), sessionCookieName)
	if err != nil {
		return err
	}
	session.Options.MaxAge = -1
	return session.Save(c.Request(), c.Response())
}
