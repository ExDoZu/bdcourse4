package server

import (
	"net/http"
	"strconv"
)

func (s *Server) cancelTourBook() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		tourBookID, err := strconv.Atoi(params.Get("tour_id"))

		if err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusBadRequest)
			return
		}

		var touristID int

		err = s.DB.QueryRow("SELECT id FROM tourist WHERE login = $1", r.Header.Get("login")).Scan(&touristID)
		if err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusBadRequest)
			return
		}

		_, err = s.DB.Exec("DELETE FROM tourist_tour WHERE tour_id = $1 and tourist_id = $2", tourBookID, touristID)
		if err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(errMsg("")))
	})
}
