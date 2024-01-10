package server

import (
	"bytes"
	"fmt"
	"io"

	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s *Server) authMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			login := r.Header.Get("login")
			password := r.Header.Get("password")

			var hashedPassword []byte
			if err := s.DB.QueryRow("SELECT password_hash FROM tourist WHERE login = $1", login).Scan(&hashedPassword); err != nil {
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}

			if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)); err != nil {
				http.Error(w, "", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

type responseWriterWrapper struct {
	responseWriter http.ResponseWriter
	status         int
	wroteHeader    bool
}

func wrap(rw http.ResponseWriter) *responseWriterWrapper {
	return &responseWriterWrapper{
		responseWriter: rw,
	}
}

func (rw *responseWriterWrapper) Status() int {
	return rw.status
}

func (rw *responseWriterWrapper) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}
	rw.status = code
	rw.responseWriter.WriteHeader(code)
	rw.wroteHeader = true
}

func (rw *responseWriterWrapper) Header() http.Header {
	return rw.responseWriter.Header()
}
func (rw *responseWriterWrapper) Write(b []byte) (int, error) {

	if !rw.wroteHeader {
		rw.status = http.StatusOK
	}
	return rw.responseWriter.Write(b)
}

func traceMsg(r *http.Request) string {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Sprintf("Headers: %s\nBody was not read: %s", r.Header, err)
	}
	r.Body = io.NopCloser(bytes.NewBuffer(data))
	return fmt.Sprintf("Headers: %s\nBody: %s", r.Header, string(data))
}

func (s *Server) logMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wrapper := wrap(w)
			traceString := traceMsg(r)
			start := time.Now()
			next.ServeHTTP(wrapper, r)
			executrionTime := time.Since(start)
			logMsg := fmt.Sprintf("Method: %s, URL: %s, Status: %d, Duration: %s, IP: %s", r.Method, r.URL, wrapper.Status(), executrionTime, strings.Split(r.RemoteAddr, ":")[0])
			if wrapper.Status() < 300 {
				log.Println(logMsg)
			} else if wrapper.Status() < 500 {
				log.Println(logMsg)
			} else {
				log.Println(logMsg)
			}
			log.Println(traceString)
			log.Println("Response Headers: ", wrapper.responseWriter.Header())
		})
	}
}

func (s *Server) addRemoteAccessMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
