package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuth(t *testing.T) {
	dbconfig := "host=localhost port=5432 user=exdo password=qaz dbname=exdodb sslmode=disable"
	server, err := New(dbconfig)
	if err != nil {
		t.Fatal(err)
	}

	// test register

	// var request struct {
	// 	FirstName     string `json:"first_name"`
	// 	LastName      string `json:"last_name"`
	// 	Age           int    `json:"age"`
	// 	Gender        bool   `json:"gender"`
	// 	CitizenshipID string `json:"citizenship_id"`
	// 	Login         string `json:"login"`
	// 	Password      string `json:"password"`
	// }

	requestData := `{
		"first_name": "test",
		"last_name": "test",
		"age": 18,
		"gender": true,
		"citizenship_id": "1",
		"login": "test",
		"password": "lkjhkj"
	}`
	reader := strings.NewReader(requestData)
	req := httptest.NewRequest("POST", "/register", reader)
	resp := httptest.NewRecorder()
	server.Router.ServeHTTP(resp, req)
	if resp.Code != 200 {
		t.Fatal("register failed", resp.Code, resp.Body.String())
	}
	// test login
	req = httptest.NewRequest("GET", "/login", nil)
	req.Header.Set("login", "test")
	req.Header.Set("password", "lkjhkj")
	resp = httptest.NewRecorder()
	server.authMiddleware()(server.Router).ServeHTTP(resp, req)
	if resp.Code == http.StatusUnauthorized || resp.Code == http.StatusBadRequest {
		t.Fatal("login failed", resp.Code, resp.Body.String())
	}
	if _, err := server.DB.Exec("DELETE FROM tourist WHERE login = 'test'"); err != nil {
		t.Fatal(err)
	}

}
