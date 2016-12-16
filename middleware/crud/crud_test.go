package crud_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/avidreder/monmach-api/middleware/crud"
	"github.com/avidreder/monmach-api/test/mocks"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type GenreTestSuite struct {
	suite.Suite
	c   echo.Context
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
	f.Set("Created", "")
	f.Set("Updated", "")
	req, err := http.NewRequest(echo.POST, "/", strings.NewReader(f.Encode()))
	req.Form = f
	req.PostForm = f
	s.NoError(err)
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/crud/:table/:id")
	c.SetParamNames("table", "id")
	c.SetParamValues("genres", "1")
	mockStore := mocks.Store{}
	mockStore.On("Create", mock.Anything).Return(nil)
	c.Set("store", mockStore)
	testHandler := echo.HandlerFunc(func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})
	h := crud.Create(testHandler)
	err = h(c)
	s.NoError(err)
}

func TestGenreTestSuite(t *testing.T) {
	suite.Run(t, new(GenreTestSuite))
}
