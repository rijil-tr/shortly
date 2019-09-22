package shortener

import (
	"time"

	"github.com/go-kit/kit/log"
	"github.com/rijil-tr/shortly"
)

type loggingService struct {
	logger log.Logger
	next   Service
}

// NewLoggingService provides a new Service with logging ebabled
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

// New wraps services New method with logs
func (s *loggingService) New(url string) (r *shortly.Link, err error) {
	defer func(begin time.Time) {
		defer func(begin time.Time) {
			s.logger.Log(
				"method", "new",
				"url", url,
				"took", time.Since(begin),
				"err", err,
			)
		}(time.Now())
	}(time.Now())
	return s.next.New(url)
}

// Get wraps services Get method with logs
func (s *loggingService) Get(id string) (r *shortly.Link, err error) {
	defer func(begin time.Time) {
		defer func(begin time.Time) {
			s.logger.Log(
				"method", "get",
				"id", id,
				"took", time.Since(begin),
				"err", err,
			)
		}(time.Now())
	}(time.Now())
	return s.next.Get(id)
}

// CountVisit wraps services CountVisit method with logs
func (s *loggingService) CountVisit(id string) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "count",
			"id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.next.CountVisit(id)
}
