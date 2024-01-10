package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) tourBook() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request bookTourRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
			return
		}

		tx, err := s.DB.Begin()
		if err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
			return
		}
		defer tx.Rollback()
		var touristID int
		if err = tx.QueryRow("SELECT id FROM tourist WHERE login = $1", r.Header.Get("login")).Scan(&touristID); err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
			return
		}
		for _, room := range request.Rooms {
			_, err = tx.Exec("call book_hotel_room($1, $2, $3, $4, $5)", touristID, room.HotelID, room.RoomID, room.StartDate, room.EndDate)
			if err != nil {
				http.Error(w, errMsg("1 "+err.Error()), http.StatusInternalServerError)
				return
			}

		}

		_, err = tx.Exec("call issue_visa($1, $2)", touristID, request.TourID)
		if err != nil {
			http.Error(w, errMsg("Visa failure"), http.StatusInternalServerError)
			return
		}

		if _, err = tx.Exec("call book_tour($1, $2)", touristID, request.TourID); err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusConflict)
			return
		}

		if err = tx.Commit(); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(errMsg("")))

	})
}
