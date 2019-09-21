package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/rijil-tr/shortly"
	"github.com/rijil-tr/shortly/shortener"
)

type shorteningHandler struct {
	s      shortener.Service
	logger log.Logger
}

var home = template.Must(template.ParseFiles("home.html"))

func (h *shorteningHandler) router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", h.Index).Methods("GET")
	r.HandleFunc("/", h.New).Methods("POST")
	r.HandleFunc("/l/{id}", h.Get).Methods("GET")
	r.HandleFunc("/s/{id}", h.Status).Methods("GET")
	return r
}

// index serve the landing page
func (h *shorteningHandler) Index(w http.ResponseWriter, r *http.Request) {
	if err := home.Execute(w, nil); err != nil {
		h.logger.Log("could not render template: ", err)
	}
}

func (h *shorteningHandler) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Code  int
		Msg   string
		Link  string
		Stats string
	}

	l, err := h.s.New(r.FormValue("link"))
	if err != nil {
		data.Code = http.StatusBadRequest
		data.Msg = "the given link is not a valid url"
	} else {
		data.Code = http.StatusCreated
		data.Msg = "link successfully created"
		data.Link = fmt.Sprintf("http://localhost:8080/l/%s", l.ID)
		data.Stats = fmt.Sprintf("http://localhost:8080/s/%s", l.ID)
	}
	if err := home.Execute(w, data); err != nil {
		h.logger.Log("could not render template:", err)
	}
}

func (h *shorteningHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[3:]
	l, err := h.s.Get(id)
	if err != nil {
		if err == shortly.ErrNoSuchLink {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	h.s.CountVisit(id)
	fmt.Fprintf(w, "<p>redirecting to %s...</p>", l.URL)
	fmt.Fprintf(w, "<script>setTimeout(function() { window.location = '%s'}, 1000)</script>", l.URL)

}

// Status tell the number of hit on a short URL
func (h *shorteningHandler) Status(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[3:]
	l, err := h.s.Get(id)
	if err != nil {
		if err == shortly.ErrNoSuchLink {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if err := json.NewEncoder(w).Encode(l); err != nil {
		h.logger.Log("could not encode link information")
	}
}
