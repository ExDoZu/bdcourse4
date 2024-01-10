package server

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (s *Server) getFreeRoomsInTour() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		tourIDstr := params.Get("tour_id")
		if tourIDstr == "" {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		tourID, err := strconv.Atoi(tourIDstr)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		// get trips in tour
		var trips []tripSimple
		rows, err := s.DB.Query("SELECT src_country_id, dst_country_id, date(datetime_start), date(datetime_end) FROM trip WHERE tour_id = $1 order by datetime_start ", tourID)
		if err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		for rows.Next() {
			var tripObj tripSimple
			err := rows.Scan(&tripObj.From, &tripObj.To, &tripObj.Start, &tripObj.End)
			if err != nil {
				http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
				return
			}
			trips = append(trips, tripObj)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
			return
		}
		rows.Close()

		// get free rooms in between trips
		var rooms []room = make([]room, 0)
		for ind, trip := range trips[:len(trips)-1] {
			if trip.To != trips[ind+1].From {
				continue
			}
			// get country between trips
			countryID, err := strconv.Atoi(trip.To)
			if err != nil {
				http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
				return
			}
			// get hotels in country
			var hotels []hotel
			rows, err := s.DB.Query("SELECT hotel_id, name, address, rating FROM hotel WHERE country_id = $1", countryID)
			if err != nil {
				http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
				return
			}
			defer rows.Close()
			for rows.Next() {
				var hotelObj hotel
				err := rows.Scan(&hotelObj.HotelID, &hotelObj.HotelName, &hotelObj.HotelAddress, &hotelObj.Rating)
				if err != nil {
					http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
					return
				}
				hotels = append(hotels, hotelObj)
			}
			if err := rows.Err(); err != nil {
				http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
				return
			}
			rows.Close()
			// get free rooms in hotels
			for _, hotel := range hotels {
				rows, err := s.DB.Query("SELECT * from get_free_rooms_by_hotel($1, $2, $3)", hotel.HotelID, trip.End, trips[ind+1].Start)
				if err != nil {
					http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
					return
				}
				defer rows.Close()
				for rows.Next() {
					var roomObj room
					err := rows.Scan(&roomObj.RoomID, &roomObj.Places, &roomObj.Price)
					if err != nil {
						http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
						return
					}
					roomObj.StartDate = trip.End
					roomObj.EndDate = trips[ind+1].Start
					roomObj.HotelID = hotel.HotelID
					roomObj.HotelName = hotel.HotelName
					roomObj.HotelAddress = hotel.HotelAddress
					roomObj.Rating = hotel.Rating
					roomObj.CountryID = countryID
					rooms = append(rooms, roomObj)
				}
				if err := rows.Err(); err != nil {
					http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
					return
				}
				rows.Close()
			}
		}
		_ = json.NewEncoder(w).Encode(rooms)

	})
}