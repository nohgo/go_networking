package api_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
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

func sendRequest(r *http.Request, fn func(w http.ResponseWriter, r *http.Request)) (*httptest.ResponseRecorder, error) {
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(fn)
	handler(w, r)
	if w.Result().StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Error code was %v and error body was: %v", w.Body, w.Result().StatusCode))
	}

	return w, nil
}

func TestMain(m *testing.M) {
	db.CreatePool()
	defer db.ClosePool()
	code := m.Run()
	os.Exit(code)
}

func TestFullProgram(t *testing.T) {

	for _, val := range mockcars {
		body, err := json.marshal(val)
		if err != nil {
			t.fatalf("json marshalling adding cars returned an error")
		}

		r, err := http.newrequest("post", "/api/cars", bytes.newbuffer(body))
		r.header.add("authorization", "bearer "+token)
		sendrequest(t, "post car", r, auth.protectedmiddle(api.postcar))
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
	carId := cars[0].Id
	r, err = http.NewRequest("DELETE", "/api/cars", bytes.NewBuffer(body))
	params := r.URL.Query()
	params.Add("id", strconv.Itoa(carId))
	r.URL.RawQuery = params.Encode()
	r.Header.Add("Authorization", "Bearer "+token)
	sendRequest(t, "Delete car", r, auth.ProtectedMiddle(api.DeleteCar))

	//Deleting User
	r, err = http.NewRequest("DELETE", "/auth", nil)
	r.Header.Add("Authorization", "Bearer "+token)
	sendRequest(t, "Delete user", r, auth.ProtectedMiddle(api.DeleteUser))
}

func TestRegister(t *testing.T) {

	userBody, err := json.Marshal(mockUser)
	if err != nil {
		t.Fatalf("json marshalling register returned an error: %v", err)
	}

	r, err := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(userBody))
	if err != nil {
		t.Fatalf("Register request creation failed: %v", err)
	}
	_, err = sendRequest(r, api.Register)

	if err != nil {
		t.Logf("Register returned an error: %v", err)
	}
}

func TestLogin(t *testing.T) {
	userBody, err := json.Marshal(mockUser)
	if err != nil {
		t.Fatalf("json marshalling register returned an error: %v", err)
	}

	r, err := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(userBody))
	loginResponse, err := sendRequest(r, api.Login)
	if err != nil {
		t.Logf("Login returned an error: %v", err)
		t.Fail()
	}
	var token string
	if err = json.Unmarshal(loginResponse.Body.Bytes(), &token); err != nil {
		t.Fatalf("Login response returned an error")
	}
	if len(token) == 0 {
		t.Fatalf("Login response returned an empty token")
	}
}
