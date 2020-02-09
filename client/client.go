package client

import (
	"mediapedia-api/database"

	"github.com/globalsign/mgo"
)

// NewSession returns a new Mongo Session.
func NewSession(server string) (database.Session, error) {
	mgoSession, err := mgo.Dial(server)
	return MongoSession{mgoSession}, err
}

// MongoSession is currently a Mongo session.
type MongoSession struct {
	*mgo.Session
}

// DB shadows *mgo.DB to returns a DataLayer interface instead of *mgo.Database.
func (s MongoSession) DB(name string) database.DataLayer {
	return &MongoDatabase{Database: s.Session.DB(name)}
}

// MongoCollection wraps a mgo.Collection to embed methods in models.
type MongoCollection struct {
	*mgo.Collection
}

func (m MongoCollection) Find(query interface{}) database.Query {
	return &MongoQuery{
		Query: m.Collection.Find(query),
	}
}

func (m MongoCollection) FindId(query interface{}) database.Query {
	return &MongoQuery{
		Query: m.Collection.FindId(query),
	}
}

// MongoDatabase wraps a mgo.Database to embed methods in models.
type MongoDatabase struct {
	*mgo.Database
}

// C shadows *mgo.DB to returns a DataLayer interface instead of *mgo.Database.
func (d MongoDatabase) C(name string) database.Collection {
	return MongoCollection{Collection: d.Database.C(name)}
}

type MongoQuery struct {
	*mgo.Query
}

func (q MongoQuery) Apply(change mgo.Change, result interface{}) (info *mgo.ChangeInfo, err error) {
	return q.Query.Apply(change, result)
}
