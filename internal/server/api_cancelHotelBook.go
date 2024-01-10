package server

import (
	"net/http"
)

func (s *Server) cancelHotelBook() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		start_date := params.Get("start_date")

		var touristID int
		if err := s.DB.QueryRow("SELECT id FROM tourist WHERE login = $1", r.Header.Get("login")).Scan(&touristID); err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusBadRequest)
			return
		}

		if _, err := s.DB.Exec("DELETE FROM tourist_hotel WHERE start_date = $1 and tourist_id = $2", start_date, touristID); err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(errMsg("")))
	})
}
