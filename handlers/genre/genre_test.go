package genre_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
  "net/url"
  "strings"
	"testing"
  "log"

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
  f := make(url.Values)
  f.Set("UserID", "1")
  f.Set("QueueID", "1")
  f.Set("Name", "test")
  f.Set("Description", "test")
  f.Set("SeedArtists", "[]")
  f.Set("SeedTracks", "[]")
  f.Set("SeedPlaylists", "[]")
  f.Set("AvatarURL", "url")
  f.Set("TrackBlacklist", "[]")
  f.Set("TrackWhitelist", "[]")
  req, err := http.NewRequest(echo.POST, "/", strings.NewReader(f.Encode()))
  req.Form = f
  req.PostForm = f
  log.Printf("%+v",req)
  s.NoError(err)
  e := echo.New()
  rec := httptest.NewRecorder()
  c := e.NewContext(req, rec)
  mockStore := mocks.Store{}
  mockStore.On("Create",mock.Anything).Return(nil)
  c.Set("store", mockStore)
  testGenreJSON, _ := json.Marshal(genreR.Genre{
    ID: 0,
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
  err = genre.Create(c)
  s.NoError(err)
  s.Equal(http.StatusOK, rec.Code)
  s.Equal(string(testGenreJSON), rec.Body.String())
}

func TestGenreTestSuite(t *testing.T) {
	suite.Run(t, new(GenreTestSuite))
}
