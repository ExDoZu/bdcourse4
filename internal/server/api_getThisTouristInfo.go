package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) getThisTouristInfo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		login := r.Header.Get("login")
		if login == "" {
			http.Error(w, errMsg("No login"), http.StatusBadRequest)
			return
		}
		var touristInfo touristInfo
		if err := s.DB.QueryRow("select first_name, last_name, age, gender, citizenship_id from tourist where login = $1", login).Scan(&touristInfo.FirstName, &touristInfo.LastName, &touristInfo.Age, &touristInfo.Gender, &touristInfo.CitizenshipID); err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(touristInfo); err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
			return
		}
	})
}