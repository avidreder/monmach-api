package crud_test

import (
	"errors"
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
		c.SetPath("/crud/:table/new")
		c.SetParamNames("table")
		c.SetParamValues(testCase.tableName)
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
		c.SetPath("/crud/:table/new")
		c.SetParamNames("table")
		c.SetParamValues(testCase.tableName)
		mockStore := mocks.Store{}
		mockStore.On("Create", mock.Anything).Return(nil)
		c.Set("store", mockStore)
		h := crud.Create(testHandler)
		err = h(c)
		s.Error(err)
	}
}

func (s *CRUDTestSuite) TestCreateBadType() {
	testHandler := echo.HandlerFunc(func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})
	testCases := []struct {
		tableName string
		model     interface{}
	}{
		{"genr", genreR.Genre{}},
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
		c.SetPath("/crud/:table/new")
		c.SetParamNames("table")
		c.SetParamValues(testCase.tableName)
		mockStore := mocks.Store{}
		mockStore.On("Create", mock.Anything).Return(nil)
		c.Set("store", mockStore)
		h := crud.Create(testHandler)
		err = h(c)
		s.Error(err)
	}
}

func (s *CRUDTestSuite) TestCreateStoreError() {
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
		c.SetPath("/crud/:table/new")
		c.SetParamNames("table")
		c.SetParamValues(testCase.tableName)
		mockStore := mocks.Store{}
		mockStore.On("Create", mock.Anything).Return(errors.New("store error"))
		c.Set("store", mockStore)
		h := crud.Create(testHandler)
		err = h(c)
		s.Error(err)
	}
}

func (s *CRUDTestSuite) TestUpdateOK() {
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
		req, err := http.NewRequest(echo.PUT, "/", strings.NewReader(f.Encode()))
		req.Form = f
		req.PostForm = f
		s.NoError(err)
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/crud/:table/:id")
		c.SetParamNames("table", "id")
		c.SetParamValues(testCase.tableName, "507f1f77bcf86cd799439011")
		mockStore := mocks.Store{}
		mockStore.On("UpdateByKey", mock.Anything).Return(nil)
		c.Set("store", mockStore)
		h := crud.Update(testHandler)
		err = h(c)
		s.NoError(err)
	}
}

func (s *CRUDTestSuite) TestUpdateBadType() {
	testHandler := echo.HandlerFunc(func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})
	testCases := []struct {
		tableName string
		model     interface{}
	}{
		{"genr", genreR.Genre{}},
	}
	for _, testCase := range testCases {
		f := make(url.Values)
		structMap := structs.Map(testCase.model)
		for k := range structMap {
			f.Set(k, "")
		}
		req, err := http.NewRequest(echo.PUT, "/", strings.NewReader(f.Encode()))
		req.Form = f
		req.PostForm = f
		s.NoError(err)
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/crud/:table/:id")
		c.SetParamNames("table", "id")
		c.SetParamValues(testCase.tableName, "507f1f77bcf86cd799439011")
		mockStore := mocks.Store{}
		mockStore.On("UpdateByKey", mock.Anything).Return(nil)
		c.Set("store", mockStore)
		h := crud.Update(testHandler)
		err = h(c)
		s.Error(err)
	}
}

func (s *CRUDTestSuite) TestUpdateBadID() {
	testHandler := echo.HandlerFunc(func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})
	testCases := []struct {
		tableName string
		model     interface{}
	}{
		{"genre", genreR.Genre{}},
	}
	for _, testCase := range testCases {
		f := make(url.Values)
		structMap := structs.Map(testCase.model)
		for k := range structMap {
			f.Set(k, "")
		}
		req, err := http.NewRequest(echo.PUT, "/", strings.NewReader(f.Encode()))
		req.Form = f
		req.PostForm = f
		s.NoError(err)
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/crud/:table/:id")
		c.SetParamNames("table", "id")
		c.SetParamValues(testCase.tableName, "bad")
		mockStore := mocks.Store{}
		mockStore.On("UpdateByKey", mock.Anything).Return(nil)
		c.Set("store", mockStore)
		h := crud.Update(testHandler)
		err = h(c)
		s.Error(err)
	}
}

func (s *CRUDTestSuite) TestUpdateStoreError() {
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
		req, err := http.NewRequest(echo.PUT, "/", strings.NewReader(f.Encode()))
		req.Form = f
		req.PostForm = f
		s.NoError(err)
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/crud/:table/:id")
		c.SetParamNames("table", "id")
		c.SetParamValues(testCase.tableName, "507f1f77bcf86cd799439011")
		mockStore := mocks.Store{}
		mockStore.On("UpdateByKey", mock.Anything).Return(errors.New("store error"))
		c.Set("store", mockStore)
		h := crud.Update(testHandler)
		err = h(c)
		s.Error(err)
	}
}

func (s *CRUDTestSuite) TestGetOK() {
	testHandler := echo.HandlerFunc(func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})
	testCases := []string{
		"genres",
		"playlists",
		"users",
		"queues",
		"tracks",
	}
	for _, testCase := range testCases {
		req, err := http.NewRequest(echo.GET, "/", nil)
		s.NoError(err)
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/crud/:table/:id")
		c.SetParamNames("table", "id")
		c.SetParamValues(testCase, "507f1f77bcf86cd799439011")
		mockStore := mocks.Store{}
		mockStore.On("GetByKey", mock.Anything).Return(nil)
		c.Set("store", mockStore)
		h := crud.Get(testHandler)
		err = h(c)
		s.NoError(err)
	}
}

func (s *CRUDTestSuite) TestGetBadType() {
	testHandler := echo.HandlerFunc(func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})
	testCases := []string{
		"genr",
	}
	for _, testCase := range testCases {
		req, err := http.NewRequest(echo.GET, "/", nil)
		s.NoError(err)
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/crud/:table/:id")
		c.SetParamNames("table", "id")
		c.SetParamValues(testCase, "507f1f77bcf86cd799439011")
		mockStore := mocks.Store{}
		mockStore.On("GetByKey", mock.Anything).Return(nil)
		c.Set("store", mockStore)
		h := crud.Get(testHandler)
		err = h(c)
		s.Error(err)
	}
}

func (s *CRUDTestSuite) TestGetBadID() {
	testHandler := echo.HandlerFunc(func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})
	testCases := []string{
		"genres",
	}
	for _, testCase := range testCases {
		req, err := http.NewRequest(echo.GET, "/", nil)
		s.NoError(err)
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/crud/:table/:id")
		c.SetParamNames("table", "id")
		c.SetParamValues(testCase, "bad")
		mockStore := mocks.Store{}
		mockStore.On("GetByKey", mock.Anything).Return(nil)
		c.Set("store", mockStore)
		h := crud.Get(testHandler)
		err = h(c)
		s.Error(err)
	}
}

func (s *CRUDTestSuite) TestGetStoreError() {
	testHandler := echo.HandlerFunc(func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})
	testCases := []string{
		"genres",
		"playlists",
		"users",
		"queues",
		"tracks",
	}
	for _, testCase := range testCases {
		req, err := http.NewRequest(echo.GET, "/", nil)
		s.NoError(err)
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/crud/:table/:id")
		c.SetParamNames("table", "id")
		c.SetParamValues(testCase, "507f1f77bcf86cd799439011")
		mockStore := mocks.Store{}
		mockStore.On("GetByKey", mock.Anything).Return(errors.New("store error"))
		c.Set("store", mockStore)
		h := crud.Get(testHandler)
		err = h(c)
		s.Error(err)
	}
}

func (s *CRUDTestSuite) TestGetAllOK() {
	testHandler := echo.HandlerFunc(func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})
	testCases := []string{
		"genres",
		"playlists",
		"users",
		"queues",
		"tracks",
	}
	for _, testCase := range testCases {
		req, err := http.NewRequest(echo.GET, "/", nil)
		s.NoError(err)
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/crud/:table/all")
		c.SetParamNames("table")
		c.SetParamValues(testCase)
		mockStore := mocks.Store{}
		mockStore.On("GetAll", mock.Anything).Return(nil)
		c.Set("store", mockStore)
		h := crud.GetAll(testHandler)
		err = h(c)
		s.NoError(err)
	}
}

func (s *CRUDTestSuite) TestGetAllStoreError() {
	testHandler := echo.HandlerFunc(func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})
	testCases := []string{
		"genres",
		"playlists",
		"users",
		"queues",
		"tracks",
	}
	for _, testCase := range testCases {
		req, err := http.NewRequest(echo.GET, "/", nil)
		s.NoError(err)
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/crud/:table/all")
		c.SetParamNames("table")
		c.SetParamValues(testCase)
		mockStore := mocks.Store{}
		mockStore.On("GetAll", mock.Anything).Return(errors.New("store error"))
		c.Set("store", mockStore)
		h := crud.GetAll(testHandler)
		err = h(c)
		s.Error(err)
	}
}

func (s *CRUDTestSuite) TestGetAllBadType() {
	testHandler := echo.HandlerFunc(func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})
	testCases := []string{
		"genr",
	}
	for _, testCase := range testCases {
		req, err := http.NewRequest(echo.GET, "/", nil)
		s.NoError(err)
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/crud/:table/all")
		c.SetParamNames("table")
		c.SetParamValues(testCase)
		mockStore := mocks.Store{}
		mockStore.On("GetAll", mock.Anything).Return(nil)
		c.Set("store", mockStore)
		h := crud.GetAll(testHandler)
		err = h(c)
		s.Error(err)
	}
}

func (s *CRUDTestSuite) TestDeleteOK() {
	testHandler := echo.HandlerFunc(func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})
	req, err := http.NewRequest(echo.DELETE, "/", nil)
	s.NoError(err)
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/crud/:table/:id")
	c.SetParamNames("table", "id")
	c.SetParamValues("genres", "507f1f77bcf86cd799439011")
	mockStore := mocks.Store{}
	mockStore.On("DeleteByKey", mock.Anything).Return(nil)
	c.Set("store", mockStore)
	h := crud.Delete(testHandler)
	err = h(c)
	s.NoError(err)
}

func (s *CRUDTestSuite) TestDeleteBadID() {
	testHandler := echo.HandlerFunc(func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})
	req, err := http.NewRequest(echo.DELETE, "/", nil)
	s.NoError(err)
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/crud/:table/:id")
	c.SetParamNames("table", "id")
	c.SetParamValues("genres", "bad")
	mockStore := mocks.Store{}
	mockStore.On("DeleteByKey", mock.Anything).Return(nil)
	c.Set("store", mockStore)
	h := crud.Delete(testHandler)
	err = h(c)
	s.Error(err)
}

func (s *CRUDTestSuite) TestDeleteStoreError() {
	testHandler := echo.HandlerFunc(func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})
	req, err := http.NewRequest(echo.DELETE, "/", nil)
	s.NoError(err)
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/crud/:table/:id")
	c.SetParamNames("table", "id")
	c.SetParamValues("genres", "507f1f77bcf86cd799439011")
	mockStore := mocks.Store{}
	mockStore.On("DeleteByKey", mock.Anything).Return(errors.New("store error"))
	c.Set("store", mockStore)
	h := crud.Delete(testHandler)
	err = h(c)
	s.Error(err)
}

func TestCRUDTestSuite(t *testing.T) {
	suite.Run(t, new(CRUDTestSuite))
}
