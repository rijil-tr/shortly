package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/rijil-tr/shortly"
	"github.com/rijil-tr/shortly/inmem"
	"github.com/rijil-tr/shortly/mongo"
	"github.com/rijil-tr/shortly/server"
	"github.com/rijil-tr/shortly/shortener"
	"gopkg.in/mgo.v2"
)

const (
	defaultHost       = "0.0.0.0"
	defaultPort       = "8080"
	defaultMongoDBURL = "mongodb://mongodb:27017"
	defaultDBName     = "shortly"
)

var (
	addr   = envString("PORT", defaultPort)
	dburl  = envString("MONGODB_URL", defaultMongoDBURL)
	dbname = envString("DB_NAME", defaultDBName)

	httpAddr     = flag.String("http.addr", "localhost:"+addr, "HTTP listen address")
	mongoDBURL   = flag.String("db.url", dburl, "MongoDB URL")
	databaseName = flag.String("db.name", dbname, "MongoDB database name")
	inmemory     = flag.Bool("inmem", false, "use in-memory repositories")

	links  shortly.LinkRepository
	logger log.Logger
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().Unix())

	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	if *inmemory {
		links = inmem.NewInMemory()
	} else {

		session, err := mgo.Dial(*mongoDBURL)
		if err != nil {
			panic(err)
		}
		defer session.Close()

		session.SetMode(mgo.Monotonic, true)

		links, _ = mongo.NewMongoRepository(*databaseName, session)
	}

	fieldKeys := []string{"method"}
	ss := shortener.NewService(links)
	ss = shortener.NewLoggingService(log.With(logger, "component", "shortener"), ss)
	ss = shortener.NewInstrumentingService(
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
		ss,
	)

	srv := server.NewSever(ss, log.With(logger, "component", "http"))
	errs := make(chan error, 2)

	go func() {
		fmt.Println(welcomeMessage())
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, srv)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

func welcomeMessage() string {
	asciiArt :=
		`

$$$$$$\  $$\                            $$\     $$\           
$$  __$$\ $$ |                           $$ |    $$ |          
$$ /  \__|$$$$$$$\   $$$$$$\   $$$$$$\ $$$$$$\   $$ |$$\   $$\ 
\$$$$$$\  $$  __$$\ $$  __$$\ $$  __$$\\_$$  _|  $$ |$$ |  $$ |
 \____$$\ $$ |  $$ |$$ /  $$ |$$ |  \__| $$ |    $$ |$$ |  $$ |
$$\   $$ |$$ |  $$ |$$ |  $$ |$$ |       $$ |$$\ $$ |$$ |  $$ |
\$$$$$$  |$$ |  $$ |\$$$$$$  |$$ |       \$$$$  |$$ |\$$$$$$$ |
 \______/ \__|  \__| \______/ \__|        \____/ \__| \____$$ |
                                                     $$\   $$ |
                                                     \$$$$$$  |
                                                      \______/ 

`
	return asciiArt
}
