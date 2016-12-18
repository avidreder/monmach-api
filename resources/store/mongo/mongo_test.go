package mongo_test

import (
	"testing"

	"github.com/avidreder/monmach-api/resources/store/mongo"
	"github.com/avidreder/monmach-api/resources/track"

	"github.com/stretchr/testify/suite"
)

type MongoTestSuite struct {
	suite.Suite
	testStore   *mongo.Store
	model       track.Track
	pluralModel []track.Track
	collection  string
}

func (s *MongoTestSuite) SetupSuite() {
	// mockStore := dbtest.DBServer{}
	// mockStore.SetPath("/")
	// mockSession := mockStore.Session()
	// s.testStore = &mongo.Store{mockSession}
	// s.collection = "tracks"
	// s.pluralModel = []track.Track{}
	// s.model = track.Track{}
}

func (s *MongoTestSuite) TestGetAll() {
	// err := s.testStore.GetAll(s.collection, s.pluralModel)
	// s.Error(err)
}

func TestMongoTestSuite(t *testing.T) {
	suite.Run(t, new(MongoTestSuite))
}
