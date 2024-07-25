package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/nohgo/go_networking/api"
	"github.com/nohgo/go_networking/api/auth"
	db "github.com/nohgo/go_networking/api/database"
	"github.com/nohgo/go_networking/api/models"
)

var mockUser models.User = models.User{
	Username: "bob",
	Password: "hello",
}

var mockCars []models.Car = []models.Car{
	{Make: "Honda", Model: "Pilot", Year: 2019},
	{Make: "Toyota", Model: "Tundra", Year: 2020},
}

func sendRequest(t *testing.T, name string, r *http.Request, fn func(w http.ResponseWriter, r *http.Request)) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(fn)
	handler(w, r)
	if w.Result().StatusCode != 200 {
		t.Fatalf("%v fail code: %v\n %v fail body: %v", name, w.Result().StatusCode, name, w.Body)
	}

	return w
}

func TestMain(m *testing.M) {
	db.CreatePool()
	defer db.ClosePool()
	code := m.Run()
	os.Exit(code)
}

func TestFullProgram(t *testing.T) {
	// Register testing
	body, err := json.Marshal(mockUser)
	if err != nil {
		t.Fatalf("json marshalling register returned an error: %v", err)
	}

	r, err := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Register request creation failed: %v", err)
	}
	sendRequest(t, "Register", r, api.Register)

	//Login Testing
	r, err = http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(body))
	loginResponse := sendRequest(t, "Login", r, api.Login)
	var token string
	if err = json.Unmarshal(loginResponse.Body.Bytes(), &token); err != nil {
		t.Fatalf("Login response returned a nil token")
	}
	if len(token) == 0 {
		t.Fatalf("Login response returned a nil token")
	}

	// Adding Cars
	for _, val := range mockCars {
		body, err := json.Marshal(val)
		if err != nil {
			t.Fatalf("Json marshalling adding cars returned an error")
		}

		r, err := http.NewRequest("POST", "/api/cars", bytes.NewBuffer(body))
		r.Header.Add("Authorization", "Bearer "+token)
		sendRequest(t, "Post Car", r, auth.ProtectedMiddle(api.PostCar))
	}

	//Getting Cars
	r, err = http.NewRequest("GET", "/api/cars", nil)
	r.Header.Add("Authorization", "Bearer "+token)
	getAllResponse := sendRequest(t, "Get Cars", r, auth.ProtectedMiddle(api.GetAll))

	var cars []models.Car
	json.NewDecoder(getAllResponse.Body).Decode(&cars)

	if len(cars) != 2 {
		t.Fatalf("Get all returned the wrong amount of cars: %v", len(cars))
	}

	//Deleting Cars
	body, err = json.Marshal(cars[0])
	if err != nil {
		t.Fatalf("Json marshalling in delete car failed: %v", err)
	}
	r, err = http.NewRequest("DELETE", "/api/cars", bytes.NewBuffer(body))
	r.Header.Add("Authorization", "Bearer "+token)
	sendRequest(t, "Delete car", r, auth.ProtectedMiddle(api.DeleteCar))

	//Deleting User
	body, err = json.Marshal(mockUser)
	if err != nil {
		t.Fatalf("Json marshalling in delete user failed: %v", err)
	}
	r, err = http.NewRequest("DELETE", "/auth", bytes.NewBuffer(body))
	r.Header.Add("Authorization", "Bearer "+token)
	sendRequest(t, "Delete user", r, auth.ProtectedMiddle(api.DeleteCar))
}
