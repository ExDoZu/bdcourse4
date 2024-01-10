package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) removeAttractionFromWishlist() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			AttractionID int `json:"attraction_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		var touristID int
		if err := s.DB.QueryRow("SELECT id FROM tourist WHERE login = $1", r.Header.Get("login")).Scan(&touristID); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		if _, err := s.DB.Exec("DELETE FROM tourist_attraction WHERE tourist_id = $1 AND attraction_id = $2", touristID, request.AttractionID); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(errMsg("")))
	})
}