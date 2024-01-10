package server

import (
	"encoding/json"
	"net/http"
)
func (s *Server) addAttractionToWishlist() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			AttractionID int `json:"attraction_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		var touristID int
		if err := s.DB.QueryRow("SELECT id FROM tourist WHERE login = $1", r.Header.Get("login")).Scan(&touristID); err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
			return
		}

		if _, err := s.DB.Exec("INSERT INTO tourist_attraction (tourist_id, attraction_id) VALUES ($1, $2)", touristID, request.AttractionID); err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusConflict)
			return
		}
		w.Write([]byte(errMsg("")))
	})
}