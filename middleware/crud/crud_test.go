package crud_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/avidreder/monmach-api/middleware/crud"
	genreR "github.com/avidreder/monmach-api/resources/genre"
	playlistR "github.com/avidreder/monmach-api/resources/playlist"
	queueR "github.com/avidreder/monmach-api/resources/queue"
	trackR "github.com/avidreder/monmach-api/resources/track"
	userR "github.com/avidreder/monmach-api/resources/user"
	"github.com/avidreder/monmach-api/test/mocks"

	"github.com/fatih/structs"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CRUDTestSuite struct {
	suite.Suite
}

func (s *CRUDTestSuite) SetupSuite() {

}

func (s *CRUDTestSuite) SetupTest() {

}

func (s *CRUDTestSuite) TestCreateOK() {
	testHandler := echo.HandlerFunc(func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})
	testCases := []struct {
		tableName string
		model     interface{}
	}{
		{"genres", genreR.Genre{}},
		{"playlists", playlistR.Playlist{}},
		{"users", userR.User{}},
		{"queues", queueR.Queue{}},
		{"tracks", trackR.Track{}},
	}
	for _, testCase := range testCases {
		f := make(url.Values)
		structMap := structs.Map(testCase.model)
		for k := range structMap {
			f.Set(k, "")
		}
		req, err := http.NewRequest(echo.POST, "/", strings.NewReader(f.Encode()))
		req.Form = f
		req.PostForm = f
		s.NoError(err)
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/crud/:table/:id")
		c.SetParamNames("table", "id")
		c.SetParamValues(testCase.tableName, "1")
		mockStore := mocks.Store{}
		mockStore.On("Create", mock.Anything).Return(nil)
		c.Set("store", mockStore)
		h := crud.Create(testHandler)
		err = h(c)
		s.NoError(err)
	}
}

func (s *CRUDTestSuite) TestCreateMissingParam() {
	testHandler := echo.HandlerFunc(func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})
	testCases := []struct {
		tableName string
		model     interface{}
	}{
		{"genres", genreR.Genre{}},
		{"playlists", playlistR.Playlist{}},
		{"users", userR.User{}},
		{"queues", queueR.Queue{}},
		{"tracks", trackR.Track{}},
	}
	for _, testCase := range testCases {
		f := make(url.Values)
		structMap := structs.Map(testCase.model)
		for k := range structMap {
			f.Set(k, "")
		}
		f.Del("Created")
		req, err := http.NewRequest(echo.POST, "/", strings.NewReader(f.Encode()))
		req.Form = f
		req.PostForm = f
		s.NoError(err)
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/crud/:table/:id")
		c.SetParamNames("table", "id")
		c.SetParamValues(testCase.tableName, "1")
		mockStore := mocks.Store{}
		mockStore.On("Create", mock.Anything).Return(nil)
		c.Set("store", mockStore)
		h := crud.Create(testHandler)
		err = h(c)
		s.NoError(err)
	}
}

func TestCRUDTestSuite(t *testing.T) {
	suite.Run(t, new(CRUDTestSuite))
}
