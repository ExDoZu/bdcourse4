package server

import (
	"encoding/json"
	"net/http"
	"strconv"
)
func (s *Server) getTourByCountry() http.Handler {
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

		var countryName string
		err = s.DB.QueryRow("SELECT name FROM country WHERE id = $1", countryID).Scan(&countryName)
		if err != nil {
			http.Error(w, errMsg("1 "+err.Error()), http.StatusInternalServerError)
			return
		}

		var tourList []tour = make([]tour, 0)
		rows, err := s.DB.Query(`select tour.id, tour.name, tour.travel_agency_id, tour.start_date, tour.end_date, tour.satisfaction_level, tour.price, agency.name 
								from tour
								inner join agency on tour.travel_agency_id = agency.id
								where tour.id in (SELECT tour_id from get_future_tours_by_country($1))`, countryName)
		if err != nil {
			http.Error(w, errMsg("2 "+err.Error()), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		for rows.Next() {
			var tour tour
			err := rows.Scan(&tour.ID, &tour.Name, &tour.TravelAgencyID, &tour.StartDate, &tour.EndDate, &tour.SatisfactionLevel, &tour.Price, &tour.TravelAgencyName)
			if err != nil {
				http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
				return
			}

			tourList = append(tourList, tour)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
			return
		}
		if err = json.NewEncoder(w).Encode(tourList); err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
			return
		}
	})
}
