package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) getCountryList() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var countryList []country
		rows, err := s.DB.Query("SELECT * FROM country")
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		for rows.Next() {
			var country country
			err := rows.Scan(&country.ID, &country.Name, &country.Capital, &country.Population, &country.Area, &country.Currency, &country.Language, &country.Climate)
			if err != nil {
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
			countryList = append(countryList, country)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		if err = json.NewEncoder(w).Encode(countryList); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	})
}
