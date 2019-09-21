package repository

import (
	"fmt"
	"math/rand"
	"net/url"
)

type inmem struct {
	links map[string]*Link
}

func randomString() string {
	return fmt.Sprintf("%X", rand.Int63())
}

// NewInMemory returns a new in-memory Link
func NewInMemory() LinkRepository {
	return &inmem{links: make(map[string]*Link)}
}

// New creates a short url
func (mem *inmem) New(u string) (*Link, error) {
	if _, err := url.ParseRequestURI(u); err != nil {
		return nil, err
	}
	l := &Link{
		ID:  randomString(),
		URL: u,
	}
	mem.links[l.ID] = l
	return l, nil
}

func (mem *inmem) Get(id string) (*Link, error) {
	l, ok := mem.links[id]
	if !ok {
		return nil, ErrNoSuchLink
	}
	return l, nil
}

func (mem *inmem) CountVisit(id string) error {
	l, ok := mem.links[id]
	if !ok {
		return ErrNoSuchLink
	}
	l.Count++
	return nil
}
