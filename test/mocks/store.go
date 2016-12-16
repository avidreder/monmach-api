package mocks

import (
	"github.com/stretchr/testify/mock"
)

// Mock is a struct for mocking a store
type Store struct {
	mock.Mock
	Arguments []interface{}
}

// QueryBIT mocks the associated store function
func (m Store) Connect() error {
	args := m.Called()
	return args.Error(0)
}

// GetByKey mocks the associated store function
func (m Store) GetByKey(string, interface{}, string, interface{}) error {
	args := m.Called()
	return args.Error(0)
}

// UpdateByKey mocks the associated store function
func (m Store) UpdateByKey(string, map[string]interface{}, string, interface{}) error {
	args := m.Called()
	return args.Error(0)
}

// DeleteByKey mocks the associated store function
func (m Store) DeleteByKey(string, string, interface{}) error {
	args := m.Called()
	return args.Error(0)
}

// Create mocks the associated store function
func (m Store) Create(string, map[string]interface{}) error {
	args := m.Called()
	return args.Error(0)
}

// GetAll mocks the associated store function
func (m Store) GetAll(string, interface{}) error {
	args := m.Called()
	return args.Error(0)
}
