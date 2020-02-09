package database

import "github.com/globalsign/mgo"

// Session is an interface to access to the Session struct.
type Session interface {
	DB(name string) DataLayer
	Close()
}

// DataLayer is an interface to access to the database struct.
type DataLayer interface {
	C(name string) Collection
}

// Collection is an interface to access to the collection struct.
type Collection interface {
	Find(query interface{}) Query
	FindId(query interface{}) Query
}

// Query is an interface to access to the Query struct.
type Query interface {
	Apply(change mgo.Change, result interface{}) (info *mgo.ChangeInfo, err error)
}
