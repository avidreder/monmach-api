package store_test

import (
	"strings"
	"testing"

	"github.com/avidreder/monmach-api/resources/store"
	trackR "github.com/avidreder/monmach-api/resources/track"

	"github.com/fatih/structs"
	"github.com/stretchr/testify/suite"
)

type StoreTestSuite struct {
	suite.Suite
}

func (s *StoreTestSuite) TestValidateRequiredOK() {
	model := trackR.Track{}
	structMap := structs.Map(model)
	structMap["Features"] = "[]"
	structMap["Artists"] = "[]"
	_, err := store.ValidateRequired(model, structMap)
	s.NoError(err)
}

func (s *StoreTestSuite) TestValidateRequiredError() {
	model := trackR.Track{}
	structMap := structs.Map(model)
	delete(structMap, "Features")
	_, err := store.ValidateRequired(model, structMap)
	s.Error(err)
}

func (s *StoreTestSuite) TestValidateRequiredBadFloatSlice() {
	model := trackR.Track{}
	structMap := structs.Map(model)
	delete(structMap, "Features")
	structMap["Features"] = "bad"
	_, err := store.ValidateRequired(model, structMap)
	s.Error(err)
}

func (s *StoreTestSuite) TestValidateRequiredBadStringSlice() {
	model := trackR.Track{}
	structMap := structs.Map(model)
	structMap["Artists"] = "bad"
	_, err := store.ValidateRequired(model, structMap)
	s.Error(err)
}

func (s *StoreTestSuite) TestValidateRequiredEmptyFloatSlice() {
	model := trackR.Track{}
	structMap := structs.Map(model)
	delete(structMap, "Features")
	structMap["Features"] = ""
	_, err := store.ValidateRequired(model, structMap)
	s.NoError(err)
}

func (s *StoreTestSuite) TestValidateRequiredEmptyStringSlice() {
	model := trackR.Track{}
	structMap := structs.Map(model)
	structMap["Artists"] = ""
	_, err := store.ValidateRequired(model, structMap)
	s.NoError(err)
}

func (s *StoreTestSuite) TestValidateInputsOK() {
	model := trackR.Track{}
	structMap := structs.Map(model)
	expected := map[string]interface{}{}
	for k, v := range structMap {
		expected[strings.ToLower(k)] = v
	}
	result := store.ValidateInputs(model, structMap)
	s.Equal(expected, result)
}

func (s *StoreTestSuite) TestValidateInputsBadFloatSlice() {
	model := trackR.Track{}
	structMap := structs.Map(model)
	expected := map[string]interface{}{}
	for k, v := range structMap {
		expected[strings.ToLower(k)] = v
	}
	delete(expected, "features")
	structMap["Features"] = "bad"
	result := store.ValidateInputs(model, structMap)
	s.Equal(expected, result)
}

func (s *StoreTestSuite) TestValidateInputsBadStringSlice() {
	model := trackR.Track{}
	structMap := structs.Map(model)
	expected := map[string]interface{}{}
	for k, v := range structMap {
		expected[strings.ToLower(k)] = v
	}
	delete(expected, "artists")
	structMap["Artists"] = "bad"
	result := store.ValidateInputs(model, structMap)
	s.Equal(expected, result)
}

func (s *StoreTestSuite) TestValidateInputsHandlesSlice() {
	model := trackR.Track{}
	structMap := structs.Map(model)
	expected := map[string]interface{}{}
	for k, v := range structMap {
		expected[strings.ToLower(k)] = v
	}
	expected["artists"] = []string{}
	expected["features"] = []float64{}
	structMap["Artists"] = "[]"
	structMap["Features"] = "[]"
	result := store.ValidateInputs(model, structMap)
	s.Equal(expected, result)
}

func TestStoreTestSuite(t *testing.T) {
	suite.Run(t, new(StoreTestSuite))
}
