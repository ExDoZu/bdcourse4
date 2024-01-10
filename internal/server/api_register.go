package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (s *Server) register() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var request struct {
			FirstName     string `json:"first_name"`
			LastName      string `json:"last_name"`
			Age           int    `json:"age"`
			Gender        bool   `json:"gender"`
			CitizenshipID string `json:"citizenship_id"`
			Login         string `json:"login"`
			Password      string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
			return
		}

		if (request.FirstName == "") || (request.LastName == "") || (request.CitizenshipID == "") || (request.Login == "") || (request.Password == "") {

			http.Error(w, errMsg("Empty field:\n"+fmt.Sprint(request)), http.StatusBadRequest)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "2", http.StatusInternalServerError)
			return
		}

		if _, err := s.DB.Exec("insert into tourist (first_name, last_name, age, gender, citizenship_id, login, password_hash) values ($1, $2, $3, $4, $5, $6, $7)",
			request.FirstName, request.LastName, request.Age, request.Gender, request.CitizenshipID, request.Login, hash); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}