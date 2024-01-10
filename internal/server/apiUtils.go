package server

import "fmt"

func errMsg(msg string) string {
	return fmt.Sprintf(`{"error": "%s"}`, msg)
}

type country struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Capital    string  `json:"capital"`
	Population int64   `json:"population"`
	Area       float64 `json:"area"`
	Currency   string  `json:"currency"`
	Language   string  `json:"language"`
	Climate    string  `json:"climate"`
}

type tour struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	TravelAgencyID    int     `json:"travel_agency_id"`
	TravelAgencyName  string  `json:"travel_agency_name"`
	StartDate         string  `json:"start_date"`
	EndDate           string  `json:"end_date"`
	SatisfactionLevel float64 `json:"satisfaction_level"`
	Price             string  `json:"price"`
}

type trip struct {
	ID            int    `json:"id"`
	SrcCountryID  int    `json:"src_country_id"`
	DstCountryID  int    `json:"dst_country_id"`
	TransportType string `json:"transport_type"`
	DatetimeStart string `json:"datetime_start"`
	DatetimeEnd   string `json:"datetime_end"`
	Price         string `json:"price"`
	TourID        int    `json:"tour_id"`
}

type guide struct {
	ID        int      `json:"id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	Gender    bool     `json:"gender"`
	AgencyID  int      `json:"agency_id"`
	Language  []string `json:"language"`
}

type tourReview struct {
	// TouristID int `json:"tourist_id"`
	// TourID int `json:"tour_id"`
	ReviewText string `json:"review_text"`
	Rating     int    `json:"rating"`
	Datetime   string `json:"datetime"`
	// ReviewSrcID int `json:"review_src_id"`
	SrcName    string `json:"src_name"`
	SrcType    string `json:"src_type"`
	SrcAddress string `json:"src_address"`
}

type tourInfo struct {
	Trips       []trip       `json:"trips"`
	Guides      []guide      `json:"guides"`
	TourReviews []tourReview `json:"reviews"`
	NeedVisa    []int        `json:"need_visa"`
}

type bookTourRequest struct {
	TourID int `json:"tour_id"`
	Rooms  []struct {
		RoomID    int    `json:"room_id"`
		HotelID   int    `json:"hotel_id"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	} `json:"room_books"`
	MakeVisa bool `json:"make_visa"`
}

type attraction struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Description string `json:"description"`
	TicketPrice string `json:"ticket_price"`
	CountryID   int    `json:"country_id"`
}
type booksResponseRoom struct {
	RoomID       int    `json:"room_id"`
	HotelID      int    `json:"hotel_id"`
	HotelAddress string `json:"hotel_address"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	HotelName    string `json:"hotel_name"`
	Places       int    `json:"places"`
	Price        string `json:"price"`
}
type booksResponse struct {
	Tours []tour              `json:"tours"`
	Rooms []booksResponseRoom `json:"rooms"`
}

type room struct {
	RoomID       int    `json:"room_id"`
	HotelID      int    `json:"hotel_id"`
	HotelName    string `json:"hotel_name"`
	HotelAddress string `json:"hotel_address"`
	Rating       int8   `json:"rating"`
	CountryID    int    `json:"country_id"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	Places       int    `json:"places"`
	Price        string `json:"price"`
}
type hotel struct {
	HotelID      int
	HotelName    string
	HotelAddress string
	Rating       int8
}

type tripSimple struct {
	From  string
	To    string
	Start string
	End   string
}

type touristInfo struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Age           int    `json:"age"`
	Gender        bool   `json:"gender"`
	CitizenshipID int    `json:"country_id"`
}
