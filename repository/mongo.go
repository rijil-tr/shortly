package repository

import (
	"net/url"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongo struct {
	db      string
	session *mgo.Session
}

// NewMongoRepository returns a MongoDB store for storing links
func NewMongoRepository(db string, session *mgo.Session) (LinkRepository, error) {
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

func (r *mongo) New(u string) (*Link, error) {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("links")

	if _, err := url.ParseRequestURI(u); err != nil {
		return nil, err
	}
	l := &Link{
		ID:  randomString(),
		URL: u,
	}
	_, err := c.Upsert(bson.M{"id": l.ID}, bson.M{"$set": l})
	if err != nil {
		return nil, err
	}

	return l, nil
}

func (r *mongo) Get(id string) (*Link, error) {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("links")
	var result Link
	if err := c.Find(bson.M{"id": id}).One(&result); err != nil {
		if err == mgo.ErrNotFound {
			return nil, ErrNoSuchLink
		}
		return nil, err
	}

	return &result, nil
}

func (r *mongo) CountVisit(id string) error {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("links")

	var result Link
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"count": 1}},
		ReturnNew: true,
	}
	_, err := c.Find(bson.M{"id": id}).Apply(change, &result)
	if err != nil {
		return ErrNoSuchLink
	}

	return nil
}
