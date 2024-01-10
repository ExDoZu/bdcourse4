package server

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (s *Server) getTourInfo() http.Handler {
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

		var tourInfo tourInfo
		tourInfo.Trips = make([]trip, 0)
		tourInfo.Guides = make([]guide, 0)
		tourInfo.TourReviews = make([]tourReview, 0)
		tourInfo.NeedVisa = make([]int, 0)

		rows, err := s.DB.Query("SELECT * FROM trip WHERE tour_id = $1", tourID)
		if err != nil {
			http.Error(w, errMsg("3"+err.Error()), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		for rows.Next() {
			var trip trip
			err := rows.Scan(&trip.ID, &trip.SrcCountryID, &trip.DstCountryID, &trip.TransportType, &trip.DatetimeStart, &trip.DatetimeEnd, &trip.Price, &trip.TourID)
			if err != nil {
				http.Error(w, errMsg("4"+err.Error()), http.StatusInternalServerError)
				return
			}

			tourInfo.Trips = append(tourInfo.Trips, trip)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, errMsg("5"+err.Error()), http.StatusInternalServerError)
			return
		}
		rows.Close()

		rows, err = s.DB.Query("SELECT guide.id, guide.first_name, guide.last_name, guide.age, guide.gender, guide.agency_id FROM guide inner join guide_tour on guide.id = guide_tour.guide_id where guide_tour.tour_id = $1", tourID)
		if err != nil {
			http.Error(w, errMsg("6"+err.Error()), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		for rows.Next() {
			var guide guide
			guide.Language = make([]string, 0)

			err = rows.Scan(&guide.ID, &guide.FirstName, &guide.LastName, &guide.Age, &guide.Gender, &guide.AgencyID)
			if err != nil {
				http.Error(w, errMsg("7"+err.Error()), http.StatusInternalServerError)
				return
			}

			

			rowsLang, err := s.DB.Query("SELECT language FROM guide_lang WHERE guide_id = $1", guide.ID)
			if err != nil {
				http.Error(w, errMsg("1 "+err.Error()), http.StatusInternalServerError)
				return
			}

			for rowsLang.Next() {
				var lang string
				err = rowsLang.Scan(&lang)
				if err != nil {
					http.Error(w, errMsg("2 "+err.Error()), http.StatusInternalServerError)
					rowsLang.Close()
					return
				}
			
				guide.Language = append(guide.Language, lang)
			}
			if err := rowsLang.Err(); err != nil {
				http.Error(w, errMsg("8"+err.Error()), http.StatusInternalServerError)
				rowsLang.Close()
				return
			}
			rowsLang.Close()

			tourInfo.Guides = append(tourInfo.Guides, guide)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, errMsg(err.Error()), http.StatusInternalServerError)
			return
		}
		rows.Close()
		rows, err = s.DB.Query("SELECT review_text, rating, datetime, name, type, address FROM tour_review inner join review_src on tour_review.review_src_id = review_src.id where tour_id = $1", tourID)
		if err != nil {
			http.Error(w, errMsg("9"+err.Error()), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		for rows.Next() {
			var tourReview tourReview
			if err = rows.Scan(&tourReview.ReviewText, &tourReview.Rating, &tourReview.Datetime, &tourReview.SrcName, &tourReview.SrcType, &tourReview.SrcAddress); err != nil {
				http.Error(w, errMsg("10"+err.Error()), http.StatusInternalServerError)
				return
			}
			tourInfo.TourReviews = append(tourInfo.TourReviews, tourReview)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, errMsg("11"+err.Error()), http.StatusInternalServerError)
			return
		}
		rows.Close()

		var touristID int
		if err = s.DB.QueryRow("SELECT id FROM tourist WHERE login = $1", r.Header.Get("login")).Scan(&touristID); err != nil {
			http.Error(w, errMsg("12"+err.Error()), http.StatusInternalServerError)
			return
		}

		rows, err = s.DB.Query("SELECT country_id from get_info_for_visa($1, $2)", touristID, tourID)
		if err != nil {
			http.Error(w, errMsg("13"+err.Error()), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		for rows.Next() {
			var countryID int
			err = rows.Scan(&countryID)
			if err != nil {
				http.Error(w, errMsg("14"+err.Error()), http.StatusInternalServerError)
				return
			}
			tourInfo.NeedVisa = append(tourInfo.NeedVisa, countryID)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, errMsg("15"+err.Error()), http.StatusInternalServerError)
			return
		}
		rows.Close()
		if err = json.NewEncoder(w).Encode(tourInfo); err != nil {
			http.Error(w, errMsg("16"+err.Error()), http.StatusInternalServerError)
			return
		}
	})
}
