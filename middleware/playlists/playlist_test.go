package bitshows_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	bitmw "github.com/avidreder/show-hawk-server/middleware/bitshows"

	"github.com/jarcoal/httpmock"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/suite"
)

type BITShowsTestSuite struct {
	suite.Suite
	Shows  bitmw.BITShows
	Client *http.Client
}

func (s *BITShowsTestSuite) SetupSuite() {
	s.Client = &http.Client{}
	tempShows := make(bitmw.BITShows, 1)
	tempShows[0] = bitmw.BITShow{
		DateTime: "2016-05-18T18:00:00",
		ID:       1,
		Artists: []bitmw.BITArtist{
			{
				"testArtist",
				"http://google.com",
				"http://google.com",
			},
		},
		Venue: bitmw.BITVenue{
			ID:        1,
			URL:       "http://google.com",
			Name:      "testVenue",
			Latitude:  45.52647,
			Longitude: -122.634697,
			Address:   "testAddress",
		},
	}
	s.Shows = tempShows
}

func testHandler(c echo.Context) error {
	return nil
}

func (s *BITShowsTestSuite) TestCompileShowsReturnsShows() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "http://api.bandsintown.com/events/search?app_id=showhawk&date=2016-08-03&format=json&location=Portland%2COR&per_page=10&radius=10",
		httpmock.NewStringResponder(200, `[{"id": 11956320,"url": "http://www.bandsintown.com/event/11956320?app_id=showhawk","datetime": "2016-05-18T18:00:00","ticket_url": "http://www.bandsintown.com/event/11956320/buy_tickets?app_id=showhawk&came_from=233","artists": [{"name": "Malachi Graham","url": "http://www.bandsintown.com/MalachiGraham","mbid": null }],"venue": { "id": 127214, "url": "http://www.bandsintown.com/venue/127214", "name": "Laurelthirst Public House", "city": "Portland", "region": "OR", "country": "United States", "latitude": 45.52647, "longitude": -122.634697}}]`))
	httpmock.RegisterResponder("GET", "http://www.bandsintown.com/MalachiGraham",
		httpmock.NewStringResponder(200, `<html><meta content="http://google.com" property="og:image" /></html>`))
	httpmock.RegisterResponder("GET", "http://maps.googleapis.com/maps/api/geocode/json?latlng=45.52647,-122.6347",
		httpmock.NewStringResponder(200, `{"results":[{"formatted_address":"test address"}]}`))
	testRoute := "/bitshows"
	req, _ := http.NewRequest("POST", testRoute, nil)
	res := httptest.NewRecorder()
	e := echo.New()
	ctx := e.NewContext(standard.NewRequest(req, nil), standard.NewResponse(res, nil))
	_ = bitmw.CompileShows(testHandler)(ctx)
	expectedShows := `[{"date":"Wed 5/18","time":"6:00 PM","id":11956320,"ticket_url":"http://www.bandsintown.com/event/11956320/buy_tickets?app_id=showhawk\u0026came_from=233","bands":["Malachi Graham"],"venue":"Laurelthirst Public House","address":"test address","image":"http://google.com","popularity":0,"neighborhood":"","price":0}]`
	actualShows, _ := json.Marshal(ctx.Get("ShowList"))
	s.Equal(expectedShows, string(actualShows[:]))
}

func (s *BITShowsTestSuite) TestQueryBITShowsReturnsShows() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://api.bandsintown.com/events/search?app_id=showhawk&date=2016-06-09&format=json&location=Portland%2COR&per_page=10&radius=10",
		httpmock.NewStringResponder(200, `[{"id": 11956320,"url": "http://www.bandsintown.com/event/11956320?app_id=showhawk","datetime": "2016-05-18T18:00:00","ticket_url": "http://www.bandsintown.com/event/11956320/buy_tickets?app_id=showhawk&came_from=233","artists": [{"name": "Malachi Graham","url": "http://www.bandsintown.com/MalachiGraham","mbid": null }],"venue": { "id": 127214, "url": "http://www.bandsintown.com/venue/127214", "name": "Laurelthirst Public House", "city": "Portland", "region": "OR", "country": "United States", "latitude": 45.52647, "longitude": -122.634697}}]`))
	testDate := "2016-06-09"
	testShows := bitmw.QueryBIT(s.Client, testDate)

	s.Equal(1, len(testShows))
}

func (s *BITShowsTestSuite) TestGetImageURLReturnsURL() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "http://google.com",
		httpmock.NewStringResponder(200, `<html><meta content="http://google.com" property="og:image" /></html>`))
	testURL := bitmw.GetImageURL(s.Client, "http://google.com")
	expected := "http://google.com"
	s.Equal(expected, testURL)
}

func (s *BITShowsTestSuite) TestGetShowImagesReturnsURLs() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "http://google.com",
		httpmock.NewStringResponder(200, `<html><meta content="http://google.com" property="og:image" /></html>`))
	testShows := bitmw.GetShowImages(s.Client, s.Shows)
	expected := "http://google.com"
	s.Equal(expected, testShows[0].ImageURL)
}

func (s *BITShowsTestSuite) TestGetAddressReturnsAddress() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "http://maps.googleapis.com/maps/api/geocode/json?latlng=45.52647,-122.63469",
		httpmock.NewStringResponder(200, `{"results":[{"formatted_address":"test address"}]}`))
	var testLatitude, testLongitude float32
	testLatitude = 45.52647
	testLongitude = -122.63469
	testAddress := bitmw.GetAddress(s.Client, testLatitude, testLongitude)
	s.Equal("test address", testAddress)
}

func (s *BITShowsTestSuite) TestGetShowAddressesReturnsAddresses() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "http://maps.googleapis.com/maps/api/geocode/json?latlng=45.52647,-122.6347",
		httpmock.NewStringResponder(200, `{"results":[{"formatted_address":"test address"}]}`))
	testShows := bitmw.GetShowAddresses(s.Client, s.Shows)
	s.Equal("test address", testShows[0].Address)
}

func (s *BITShowsTestSuite) TestParseDateAndTimeReturnsDateAndTime() {
	testDate, testTime, _ := bitmw.ParseDateAndTime("2016-05-18T18:00:00")
	s.Equal("Wed 5/18", testDate)
	s.Equal("6:00 PM", testTime)
}

func (s *BITShowsTestSuite) TestParseShowDatesReturnsDateAndTime() {
	testShows := bitmw.ParseShowDates(s.Shows)
	s.Equal("Wed 5/18", testShows[0].Date)
	s.Equal("6:00 PM", testShows[0].Time)
}

func (s *BITShowsTestSuite) TestMapBITShowsToShowsReturnsShows() {
	testShows, _ := bitmw.MapBITShowsToShows(s.Shows)
	s.IsType(bitmw.Shows{}, testShows)
}

func TestBITShowsTestSuite(t *testing.T) {
	suite.Run(t, new(BITShowsTestSuite))
}
