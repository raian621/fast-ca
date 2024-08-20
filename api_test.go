package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/raian621/fast-ca/api"
	"github.com/raian621/fast-ca/database"
	"github.com/raian621/fast-ca/util"
	"github.com/stretchr/testify/assert"
)

var testClientDir = "./client/dist"

func TestSignIn(t *testing.T) {
  absClientDir, err := util.RelativeToAbsolutePath(testClientDir)
  if err != nil {
    t.Fatal(err)
  }
  username := "user1"
  password := "password123"
  email := "user1@localhost"
  salt, err := util.GenerateSalt()
  assert.NoError(t, err)
  passhash := util.HashPassword(password, salt) 

  id, err := database.CreateUser(username, email, passhash)
  if err != nil {
    t.Fatal(err)
  }
  defer func() {
    if err := database.DeleteUser(id); err != nil {
      t.Fatal(err)
    }
  }()

  userCreds := api.UserCredentials{
    Username: username,
    Password: password,
  }
  userCredsData := bytes.Buffer{} 
  if err := json.NewEncoder(&userCredsData).Encode(userCreds); err != nil {
    t.Fatal(err)
  }

  server := NewServer(absClientDir)
  req := httptest.NewRequest("POST", "/api/v1/signin", &userCredsData)
  rec := httptest.NewRecorder()
  ctx := server.e.NewContext(req, rec)  

  if err := server.PostSignin(ctx); err != nil {
    t.Fatal(err)
  }

  assert.Equal(t, 200, rec.Code)
  assert.Contains(t, rec.Header(), "Set-Cookie")

  userCreds.Username = "nottheusername"
  userCredsData.Reset() 
  if err := json.NewEncoder(&userCredsData).Encode(userCreds); err != nil {
    t.Fatal(err)
  }
  t.Log(userCredsData.String())
  req = httptest.NewRequest("POST", "/api/v1/signin", &userCredsData)
  rec = httptest.NewRecorder()
  ctx = server.e.NewContext(req, rec)  

  assert.Equal(t, 400, rec.Code)
  assert.NotContains(t, rec.Header(), "Set-Cookie")

  userCreds.Username = username
  userCreds.Password = "nothepassword"
  userCredsData.Reset()
  if err := json.NewEncoder(&userCredsData).Encode(userCreds); err != nil {
    t.Fatal(err)
  }
  req = httptest.NewRequest("POST", "/api/v1/signin", &userCredsData)
  rec = httptest.NewRecorder()
  ctx = server.e.NewContext(req, rec)  

  assert.Equal(t, 400, rec.Code)
  assert.NotContains(t, rec.Header(), "Set-Cookie")
}
