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

// Get mocks the associated store function
func (m Store) Get(interface{}) error {
	args := m.Called()
	return args.Error(0)
}

// Update mocks the associated store function
func (m Store) Update(string, int64, map[string]interface{}) error {
	args := m.Called()
	return args.Error(0)
}

// Delete mocks the associated store function
func (m Store) Delete(interface{}) error {
	args := m.Called()
	return args.Error(0)
}

// Create mocks the associated store function
func (m Store) Create(interface{}) error {
	args := m.Called()
	return args.Error(0)
}

// GetAll mocks the associated store function
func (m Store) GetAll(interface{}, string) error {
	args := m.Called()
	return args.Error(0)
}
