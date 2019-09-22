package shortener

import "github.com/rijil-tr/shortly"

// A Service provides all the operations to create and store links.
type Service interface {
	New(url string) (*shortly.Link, error)
	Get(id string) (*shortly.Link, error)
	CountVisit(id string) error
}

// service implements Service functionalities
type service struct {
	store shortly.LinkRepository
}

// NewService creates a new service with a backing store
func NewService(store shortly.LinkRepository) Service {
	return &service{
		store: store,
	}
}

// New creates a short url and stores in it in the backing store
func (s *service) New(url string) (*shortly.Link, error) {
	return s.store.New(url)
}

// Get retrives the long url given the short one
func (s *service) Get(id string) (*shortly.Link, error) {
	return s.store.Get(id)
}

// CountVisit update the link visit count
func (s *service) CountVisit(id string) error {
	return s.store.CountVisit(id)
}
