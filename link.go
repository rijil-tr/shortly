package shortly

import "errors"

// A Link contains the information related to a shorten link.
type Link struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Count int64  `json:"count"`
}

// ErrNoSuchLink is thrown when link not found
var ErrNoSuchLink = errors.New("no such link")

// A Service provides all the operations to create and store links.
type LinkRepository interface {
	New(url string) (*Link, error)
	Get(id string) (*Link, error)
	CountVisit(id string) error
}
