package server

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (s *Server) getAttractions() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		countryIDstr := params.Get("country_id")
		if countryIDstr == "" {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		countryID, err := strconv.Atoi(countryIDstr)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		rows, err := s.DB.Query("SELECT * FROM attraction WHERE country_id = $1", countryID)
		if err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		var attractionList []attraction
		for rows.Next() {
			var attraction attraction
			err := rows.Scan(&attraction.ID, &attraction.Name, &attraction.Address, &attraction.Description, &attraction.TicketPrice, &attraction.CountryID)
			if err != nil {
				http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
				return
			}
			attractionList = append(attractionList, attraction)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
			return
		}
		if err = json.NewEncoder(w).Encode(attractionList); err != nil {
			http.Error(w, ":-(", http.StatusInternalServerError)
			return
		}
	})
}