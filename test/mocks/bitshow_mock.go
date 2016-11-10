package mock

import (
	"net/http"

	"github.com/avidreder/show-hawk-server/middleware/bitshows"
	"github.com/stretchr/testify/mock"
)

// Mock is a struct for mocking bitshow responses
type Mock struct {
	mock.Mock
	Arguments []interface{}
}

// QueryBIT mocks the associated bitshows function
func (m *Mock) QueryBIT(*http.Client, string) bitshows.BITShows {
	args := m.Called()
	return args.Get(0).(bitshows.BITShows)
}

// GetImageURL mocks the associated bitshows function
func (m *Mock) GetImageURL(*http.Client, string) string {
	args := m.Called()
	return args.Get(0).(string)
}

// GetAddress mocks the associated bitshows function
func (m *Mock) GetAddress(*http.Client, float32, float32) string {
	args := m.Called()
	return args.Get(0).(string)
}
