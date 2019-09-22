package shortener

import (
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/rijil-tr/shortly"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           Service
}

// NewInstrumentingService will create a new service by wrapping it over request latency
func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		next:           s,
	}
}

// New gives metrics of POST / endpoint
func (s *instrumentingService) New(url string) (*shortly.Link, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "new").Add(1)
		s.requestLatency.With("method", "new").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.New(url)
}

// Get gives metrics of GET /l/{id} endpoint
func (s *instrumentingService) Get(id string) (*shortly.Link, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "get").Add(1)
		s.requestLatency.With("method", "get").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.Get(id)
}

// Get gives metrics of GET /s/{id} endpoint
func (s *instrumentingService) CountVisit(id string) error {
	defer func(begin time.Time) {
		s.requestCount.With("method", "visit").Add(1)
		s.requestLatency.With("method", "visit").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.CountVisit(id)
}
