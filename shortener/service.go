package shortener

import "github.com/rijil-tr/shortly"

// A Service provides all the operations to create and store links.
type Service interface {
	New(url string) (*shortly.Link, error)
	Get(id string) (*shortly.Link, error)
	CountVisit(id string) error
}

type service struct {
	store shortly.LinkRepository
}

func NewService(store shortly.LinkRepository) Service {
	return &service{
		store: store,
	}
}

func (s *service) New(url string) (*shortly.Link, error) {
	return s.store.New(url)
}
func (s *service) Get(id string) (*shortly.Link, error) {
	return s.store.Get(id)
}
func (s *service) CountVisit(id string) error {
	return s.store.CountVisit(id)
}
