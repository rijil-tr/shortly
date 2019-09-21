package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"text/template"
	"time"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/gorilla/mux"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rijil-tr/shortly/repository"
	"gopkg.in/mgo.v2"
)

const (
	defaultHost       = "localhost"
	defaultPort       = "8080"
	defaultMongoDBURL = "127.0.0.1"
	defaultDBName     = "shortly"
)

var (
	home   = template.Must(template.ParseFiles("home.html"))
	addr   = envString("PORT", defaultPort)
	dburl  = envString("MONGODB_URL", defaultMongoDBURL)
	dbname = envString("DB_NAME", defaultDBName)

	httpAddr     = flag.String("http.addr", "localhost:"+addr, "HTTP listen address")
	mongoDBURL   = flag.String("db.url", dburl, "MongoDB URL")
	databaseName = flag.String("db.name", dbname, "MongoDB database name")
	inmemory     = flag.Bool("inmem", false, "use in-memory repositories")
	// ctx          = context.Background()

	links  repository.LinkRepository
	logger log.Logger
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().Unix())

	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	if *inmemory {
		links = repository.NewInMemory()
	} else {

		session, err := mgo.Dial(*mongoDBURL)
		if err != nil {
			panic(err)
		}
		defer session.Close()

		session.SetMode(mgo.Monotonic, true)

		links, _ = repository.NewMongoRepository(*databaseName, session)
	}
	fieldKeys := []string{"method"}
	links = NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "url_shortening_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "url_shortening_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		links,
	)
	links = NewLoggingService(logger, links)
	r := mux.NewRouter()
	r.Use(accessControl)
	r.HandleFunc("/health", health)
	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/", newURL).Methods("POST")
	r.HandleFunc("/l/{id}", getURL).Methods("GET")
	r.HandleFunc("/s/{id}", getStatus).Methods("GET")
	r.Handle("/metrics", promhttp.Handler()).Methods("GET")
	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, r)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}

// Home serve the landing page
func index(w http.ResponseWriter, r *http.Request) {
	if err := home.Execute(w, nil); err != nil {
		logger.Log("could not render template: ", err)
	}
}

func newURL(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Code  int
		Msg   string
		Link  string
		Stats string
	}

	l, err := links.New(r.FormValue("link"))
	if err != nil {
		data.Code = http.StatusBadRequest
		data.Msg = "the given link is not a valid url"
	} else {
		data.Code = http.StatusCreated
		data.Msg = "link successfully created"
		data.Link = fmt.Sprintf("http://%s/l/%s", *httpAddr, l.ID)
		data.Stats = fmt.Sprintf("http://%s/s/%s", *httpAddr, l.ID)
	}
	if err := home.Execute(w, data); err != nil {
		logger.Log("could not render template:", err)
	}
}

func getURL(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[3:]
	l, err := links.Get(id)
	if err != nil {
		if err == repository.ErrNoSuchLink {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	links.CountVisit(id)
	fmt.Fprintf(w, "<p>redirecting to %s...</p>", l.URL)
	fmt.Fprintf(w, "<script>setTimeout(function() { window.location = '%s'}, 1000)</script>", l.URL)

}

// GetStatus tell the number of hit on a short URL
func getStatus(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[3:]
	l, err := links.Get(id)
	if err != nil {
		if err == repository.ErrNoSuchLink {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if err := json.NewEncoder(w).Encode(l); err != nil {
		logger.Log("could not encode link information")
	}
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
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
