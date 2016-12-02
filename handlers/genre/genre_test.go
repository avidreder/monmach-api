package genre_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/avidreder/monmach-api/handlers/genre"
  genreR "github.com/avidreder/monmach-api/resources/genre"
  "github.com/avidreder/monmach-api/test/mocks"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
  "github.com/stretchr/testify/mock"
)

type GenreTestSuite struct {
	suite.Suite
	c echo.Context
  req *http.Request
  rec *httptest.ResponseRecorder
}

func (s *GenreTestSuite) SetupSuite() {

}

func (s *GenreTestSuite) SetupTest() {
  e := echo.New()
	s.req = new(http.Request)
	s.rec = httptest.NewRecorder()
  s.c = e.NewContext(s.req, s.rec)
}

func (s *GenreTestSuite) TestCreateOK() {
  mockStore := mocks.Store{}
  mockStore.On("Create",mock.Anything).Return(nil)
  s.c.Set("store", &mockStore)
  testGenreJSON, _ := json.Marshal(genreR.Genre{
    ID: 1,
  	UserID: 1,
  	QueueID:1,
  	Name: "test",
  	Description: "test",
  	SeedArtists: []int64{},
  	SeedTracks: []int64{},
  	SeedPlaylists: []int64{},
  	AvatarURL: "url",
  	TrackBlacklist: []int64{},
  	TrackWhitelist: []int64{},
  })
  err := genre.Create(s.c)
  s.NoError(err)
  s.Equal(http.StatusOK, s.rec.Code)
  s.Equal(testGenreJSON, s.rec.Body.String())
}

func TestGenreTestSuite(t *testing.T) {
	suite.Run(t, new(GenreTestSuite))
}
