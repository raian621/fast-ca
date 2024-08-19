package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key      = "SESSION_SECRET"
//	maxAge   = 86400 * 30
//	HttpOnly = true
//	Secure   = false
)

var store sessions.Store

func init() {
	store = sessions.NewCookieStore([]byte(key))
}

func AuthenticateUser(username, password string) (bool, error) {
	return false, nil
}

func AuthenticateApiKey(apiKey string) (bool, error) {
	return false, nil
}

func GetSession(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, "FASTCA_SESSION")
}
