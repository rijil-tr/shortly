package memory

import (
	"fmt"
	"math/rand"
	"net/url"

	"github.com/rijil-tr/shortly"
)

type inmem struct {
	links map[string]*shortly.Link
}

func randomString() string {
	return fmt.Sprintf("%X", rand.Int63())
}

// NewInMemory returns a new in-memory Link
func NewInMemory() shortly.LinkRepository {
	return &inmem{links: make(map[string]*shortly.Link)}
}

// New creates a short url
func (mem *inmem) New(u string) (*shortly.Link, error) {
	if _, err := url.ParseRequestURI(u); err != nil {
		return nil, err
	}
	l := &shortly.Link{
		ID:  randomString(),
		URL: u,
	}
	mem.links[l.ID] = l
	return l, nil
}

func (mem *inmem) Get(id string) (*shortly.Link, error) {
	l, ok := mem.links[id]
	if !ok {
		return nil, shortly.ErrNoSuchLink
	}
	return l, nil
}

func (mem *inmem) CountVisit(id string) error {
	l, ok := mem.links[id]
	if !ok {
		return shortly.ErrNoSuchLink
	}
	l.Count++
	return nil
}
