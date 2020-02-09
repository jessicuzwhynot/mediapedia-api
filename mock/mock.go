package mock

import (
	"mediapedia-api/database"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// MockSession satisfies Session and act as a mock of *mgo.session.
type MockSession struct{}

// NewMockSession mock NewSession.
func NewMockSession() database.Session {
	return MockSession{}
}

// Close mocks mgo.Session.Close().
func (fs MockSession) Close() {}

// DB mocks mgo.Session.DB().
func (fs MockSession) DB(name string) database.DataLayer {
	mockDatabase := MockDatabase{}
	return mockDatabase
}

// MockDatabase satisfies DataLayer and act as a mock.
type MockDatabase struct{}

// MockCollection satisfies Collection and act as a mock.
type MockCollection struct{}

// Find mock
func (fc MockCollection) Find(query interface{}) database.Query {
	return MockQuery{}
}

// FindId mock
func (fc MockCollection) FindId(query interface{}) database.Query {
	return MockQuery{}
}

// C mocks mgo.Database(name).Collection(name).
func (db MockDatabase) C(name string) database.Collection {
	mockCollection := MockCollection{}
	return mockCollection
}

type MockQuery struct{}

var ApplyResult = map[string]interface{}{
	"siteId": "fp-us",
	"NextId": 9999,
	"prefix": "FX",
}

func (mq MockQuery) Apply(change mgo.Change, result interface{}) (info *mgo.ChangeInfo, err error) {
	*result.(*bson.M) = ApplyResult
	return &mgo.ChangeInfo{}, nil
}
