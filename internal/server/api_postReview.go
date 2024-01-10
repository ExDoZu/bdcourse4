package server

import (
	"encoding/json"
	"net/http"
	"time"
)

func (s *Server) postReview() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			TourID     int    `json:"tour_id"`
			ReviewText string `json:"text"`
			Rating     int    `json:"rating"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusBadRequest)
			return
		}
		var touristID int
		if err := s.DB.QueryRow("SELECT id FROM tourist WHERE login = $1", r.Header.Get("login")).Scan(&touristID); err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
			return
		}
		if _, err := s.DB.Exec("INSERT INTO tour_review (tourist_id, tour_id, review_text, rating, datetime, review_src_id) VALUES ($1, $2, $3, $4, $5, $6)", touristID, request.TourID, request.ReviewText, request.Rating, time.Now(), 1); err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusConflict)
			return
		}
		w.Write([]byte(errMsg("")))
	})
}