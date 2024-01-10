package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"bdcourse/internal/db"
)

type Server struct {
	DB     *sql.DB
	Router *mux.Router
}

func New(dbconfig string) (*Server, error) {
	db, err := db.New(dbconfig)
	if err != nil {
		return nil, err
	}
	server := &Server{
		DB: db,
	}

	server.configureRouter()
	return server, nil
}

func (s *Server) Run(addr string) error {
	log.Println("Server is listening on", addr)
	return http.ListenAndServe(addr, s.Router)
}

func (s *Server) configureRouter() {
	s.Router = mux.NewRouter()
	s.Router.Use(s.logMiddleware(), s.addRemoteAccessMiddleware())

	s.Router.Handle("/tour/book", s.authMiddleware()(s.tourBook()))
	s.Router.Handle("/tour/postReview", s.authMiddleware()(s.postReview()))
	s.Router.Handle("/getBooks", s.authMiddleware()(s.getBooks()))
	s.Router.Handle("/attraction/addToWishlist", s.authMiddleware()(s.addAttractionToWishlist()))
	s.Router.Handle("/attraction/removeFromWishlist", s.authMiddleware()(s.removeAttractionFromWishlist()))
	s.Router.Handle("/attraction/getWishlist", s.authMiddleware()(s.getWishlist()))
	s.Router.Handle("/tour/getInfo", s.authMiddleware()(s.getTourInfo()))
	s.Router.Handle("/whoami", s.authMiddleware()(s.getThisTouristInfo()))
	s.Router.Handle("/cancelTourBook", s.authMiddleware()(s.cancelTourBook()))
	s.Router.Handle("/cancelHotelBook", s.authMiddleware()(s.cancelHotelBook()))

	s.Router.Handle("/tour/getByCountry", s.getTourByCountry())
	s.Router.Handle("/tour/freeRooms", s.getFreeRoomsInTour())
	s.Router.Handle("/getCountryList", s.getCountryList())
	s.Router.Handle("/attraction/getList", s.getAttractions())
	s.Router.Handle("/register", s.register())
}
