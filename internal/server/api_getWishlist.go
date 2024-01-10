package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) getWishlist() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var touristID int
		if err := s.DB.QueryRow("SELECT id FROM tourist WHERE login = $1", r.Header.Get("login")).Scan(&touristID); err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
			return
		}
		rows, err := s.DB.Query("SELECT attraction.id, attraction.name, attraction.address, attraction.description, attraction.ticket_price, attraction.country_id FROM attraction inner join tourist_attraction on tourist_attraction.attraction_id = id WHERE tourist_id = $1", touristID)
		if err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		var attractionList []attraction = make([]attraction, 0)
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
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	})
}
