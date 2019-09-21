package main

import (
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/rijil-tr/shortly/repository"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           repository.LinkRepository
}

func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s repository.LinkRepository) repository.LinkRepository {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		next:           s,
	}
}

func (s *instrumentingService) New(url string) (*repository.Link, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "new").Add(1)
		s.requestLatency.With("method", "new").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.New(url)
}
func (s *instrumentingService) Get(id string) (*repository.Link, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "get").Add(1)
		s.requestLatency.With("method", "get").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.Get(id)
}
func (s *instrumentingService) CountVisit(id string) error {
	defer func(begin time.Time) {
		s.requestCount.With("method", "visit").Add(1)
		s.requestLatency.With("method", "visit").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.CountVisit(id)
}
