package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) getBooks() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var touristID int
		if err := s.DB.QueryRow("SELECT id FROM tourist WHERE login = $1", r.Header.Get("login")).Scan(&touristID); err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
			return
		}

		rows, err := s.DB.Query(`SELECT tour.id, tour.name, tour.travel_agency_id, tour.start_date, tour.end_date, tour.satisfaction_level, tour.price, agency.name
								FROM tour 
								inner join tourist_tour on tour.id = tourist_tour.tour_id 
								inner join agency on tour.travel_agency_id = agency.id
								WHERE tourist_tour.tourist_id = $1`, touristID)
		if err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		var books booksResponse
		books.Tours = make([]tour, 0)
		books.Rooms = make([]booksResponseRoom, 0)

		for rows.Next() {
			var tour tour
			err := rows.Scan(&tour.ID, &tour.Name, &tour.TravelAgencyID, &tour.StartDate, &tour.EndDate, &tour.SatisfactionLevel, &tour.Price, &tour.TravelAgencyName)
			if err != nil {
				http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
				return
			}

			books.Tours = append(books.Tours, tour)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
			return
		}
		rows.Close()

		rows, err = s.DB.Query(`SELECT tourist_hotel.room_id, tourist_hotel.hotel_id, hotel.address, start_date, end_date, hotel.name, hotel_room.places_number, hotel_room.price_per_night 
								FROM tourist_hotel 
								inner join hotel on hotel.hotel_id = tourist_hotel.hotel_id 
								inner join hotel_room on hotel_room.hotel_id = tourist_hotel.hotel_id and hotel_room.room_id = tourist_hotel.room_id
								where tourist_id = $1`, touristID)
		if err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		for rows.Next() {
			var room booksResponseRoom
			err := rows.Scan(&room.RoomID, &room.HotelID, &room.HotelAddress, &room.StartDate, &room.EndDate, &room.HotelName, &room.Places, &room.Price)
			if err != nil {
				http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
				return
			}
			books.Rooms = append(books.Rooms, room)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
		}
		rows.Close()

		_ = json.NewEncoder(w).Encode(books)

	})
}