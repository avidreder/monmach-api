package mongo_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MongoTestSuite struct {
	suite.Suite
}

func (s *MongoTestSuite) TestValidateRequiredOK() {

}

func TestMongoTestSuite(t *testing.T) {
	suite.Run(t, new(MongoTestSuite))
}
