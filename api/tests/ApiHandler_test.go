package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/nohgo/go_networking/api"
	db "github.com/nohgo/go_networking/api/database"
	"github.com/nohgo/go_networking/api/models"
)

var user models.User

func TestMain(m *testing.M) {
	user = models.User{
		Username: "bob",
		Password: "hello",
	}
	db.CreatePool()
	defer db.ClosePool()
	code := m.Run()
	os.Exit(code)
}

func TestRegister(t *testing.T) {
	body, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Json marshalling returned an error: %v", err)
	}
	r, err := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Request creation failed: %v", err)
	}
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(api.Register)
	handler(w, r)
	if w.Result().StatusCode != 200 {
		t.Fatalf("Register returned a fail code: %v", w.Result().StatusCode)
	}
}
