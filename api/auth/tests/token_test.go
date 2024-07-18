package auth_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/nohgo/go_networking/api/auth"
	"github.com/nohgo/go_networking/api/models"
)

var user models.User

func TestMain(m *testing.M) {
	user = models.User{
		Username: "noh",
	}
	code := m.Run()
	os.Exit(code)
}

func TestCreateParse(t *testing.T) {
	token, err := auth.CreateToken(user.Username)

	if err != nil {
		t.Fatalf("Token creation returned an error: %v", err)
	}
	if len(token) == 0 {
		t.Fatalf("Token creation returned an empty string")
	}

	parsedName, err := auth.ParseToken(token)
	if err != nil {
		t.Fatalf("Parse token returned an error: %v, for token: %v", err, token)
	}
	if len(parsedName) == 0 {
		t.Fatalf("Parse token returned an empty string")
	}
	if parsedName != user.Username {
		t.Fatalf("Parse token returned the wrong name: %v", parsedName)
	}
}

func TestMiddleware(t *testing.T) {
	token, err := auth.CreateToken(user.Username)
	if err != nil {
		t.Fatalf("Token creation returned an error: %v", err)
	}

	req, err := http.NewRequest(http.MethodGet, "/cars", nil)
	if err != nil {
		t.Fatalf("Response creation failed: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	res := httptest.NewRecorder()
	handler := auth.ProtectedMiddle(func(w http.ResponseWriter, r *http.Request) {
	})
	handler(res, req)

	if res.Result().StatusCode != 200 {
		t.Fatalf("Middleware returned code: %v", res.Code)
	}

	if len(req.Header.Get("Authorization")) == 0 {
		t.Fatalf("Response has no authorization header")
	}

	if req.Header.Get("Authorization") != user.Username {
		t.Fatalf("Response returned the incorrect username: %v", res.Header().Get("Authorization"))
	}

}
