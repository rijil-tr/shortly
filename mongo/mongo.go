package mongo

import (
	"fmt"
	"math/rand"
	"net/url"

	"github.com/rijil-tr/shortly"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongo struct {
	db      string
	session *mgo.Session
}

// NewMongoRepository returns a MongoDB store for storing links
func NewMongoRepository(db string, session *mgo.Session) (shortly.LinkRepository, error) {
	r := &mongo{
		db:      db,
		session: session,
	}
	index := mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	sess := r.session.Copy()
	defer sess.Close()
	c := sess.DB(r.db).C("links")

	if err := c.EnsureIndex(index); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *mongo) New(u string) (*shortly.Link, error) {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("links")

	if _, err := url.ParseRequestURI(u); err != nil {
		return nil, err
	}
	l := &shortly.Link{
		ID:  randomString(),
		URL: u,
	}
	_, err := c.Upsert(bson.M{"id": l.ID}, bson.M{"$set": l})
	if err != nil {
		return nil, err
	}

	return l, nil
}

func (r *mongo) Get(id string) (*shortly.Link, error) {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("links")
	var result shortly.Link
	if err := c.Find(bson.M{"id": id}).One(&result); err != nil {
		if err == mgo.ErrNotFound {
			return nil, shortly.ErrNoSuchLink
		}
		return nil, err
	}

	return &result, nil
}

func (r *mongo) CountVisit(id string) error {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("links")

	var result shortly.Link
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"count": 1}},
		ReturnNew: true,
	}
	_, err := c.Find(bson.M{"id": id}).Apply(change, &result)
	if err != nil {
		return shortly.ErrNoSuchLink
	}

	return nil
}

func randomString() string {
	return fmt.Sprintf("%X", rand.Int63())
}
