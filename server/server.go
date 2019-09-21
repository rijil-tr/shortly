package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rijil-tr/shortly/shortener"
)

type Server struct {
	Shortener shortener.Service
	Logger    log.Logger
	router    *mux.Router
}

func NewSever(ss shortener.Service, logger log.Logger) *Server {
	s := &Server{
		Shortener: ss,
		Logger:    logger,
	}

	r := mux.NewRouter()
	r.Use(accessControl)
	r.HandleFunc("/health", health)
	r.Handle("/metrics", promhttp.Handler()).Methods("GET")
	h := shorteningHandler{
		s:      ss,
		logger: s.Logger,
	}
	r.PathPrefix("/").Handler(h.router())
	s.router = r
	return s
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
