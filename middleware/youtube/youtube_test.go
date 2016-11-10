package youtube_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/avidreder/show-hawk-server/middleware/youtube"

	"github.com/jarcoal/httpmock"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/suite"
)

type YoutubeTestSuite struct {
	suite.Suite
}

func testHandler(c echo.Context) error {
	return nil
}

func (s *YoutubeTestSuite) TestCompileVideosSetsVideos() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://www.googleapis.com/youtube/v3/search?maxResults=5&type=video&key=AIzaSyD5e2odC6GoTPF73GEF1uoN3o2OT9eCg5k&part=snippet&q=music",
		httpmock.NewStringResponder(200, `{"kind":"youtube#searchListResponse","etag":"\"I_8xdZu766_FSaexEaDXTIfEWc0/SOzT2IaFjcn5yxsN-DdBk41zd4E\"","nextPageToken":"CAEQAA","regionCode":"US","pageInfo":{"totalResults":1000000,"resultsPerPage":1},"items":[{"kind":"youtube#searchResult","etag":"\"I_8xdZu766_FSaexEaDXTIfEWc0/GEeQVBHgE63BMvLjK7YJRMv8_zc\"","id":{"kind":"youtube#video","videoId":"y882AFjrSOM"},"snippet":{"publishedAt":"2016-07-31T15:00:02.000Z","channelId":"UCoc2P5sEaBV4p173cYBzDpA","title":"HyunA(현아) - '어때? (How's this?)' Official Music Video","description":"HyunA(현아) - '어때? (How's this?)' Official Music Video.","thumbnails":{"default":{"url":"https://i.ytimg.com/vi/y882AFjrSOM/default.jpg","width":120,"height":90},"medium":{"url":"https://i.ytimg.com/vi/y882AFjrSOM/mqdefault.jpg","width":320,"height":180},"high":{"url":"https://i.ytimg.com/vi/y882AFjrSOM/hqdefault.jpg","width":480,"height":360}},"channelTitle":"HyunA 현아 (Official YouTube Channel)","liveBroadcastContent":"none"}}]}`))
	testRoute := "/youtube"
	req, _ := http.NewRequest("POST", testRoute, nil)
	res := httptest.NewRecorder()
	e := echo.New()
	ctx := e.NewContext(standard.NewRequest(req, nil), standard.NewResponse(res, nil))
	_ = youtube.CompileVideos(testHandler)(ctx)
	expectedVideos := `[{"title":"HyunA(현아) - '어때? (How's this?)' Official Music Video","id":"y882AFjrSOM"}]`
	actualVideos, _ := json.Marshal(ctx.Get("VideoList"))
	s.Equal(expectedVideos, string(actualVideos[:]))
}

func (s *YoutubeTestSuite) TestQueryYoutubeReturnsRawVideos() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://www.googleapis.com/youtube/v3/search?maxResults=5&type=video&key=AIzaSyD5e2odC6GoTPF73GEF1uoN3o2OT9eCg5k&part=snippet&q=test",
		httpmock.NewStringResponder(200, `{"kind":"youtube#searchListResponse","etag":"\"I_8xdZu766_FSaexEaDXTIfEWc0/SOzT2IaFjcn5yxsN-DdBk41zd4E\"","nextPageToken":"CAEQAA","regionCode":"US","pageInfo":{"totalResults":1000000,"resultsPerPage":1},"items":[{"kind":"youtube#searchResult","etag":"\"I_8xdZu766_FSaexEaDXTIfEWc0/GEeQVBHgE63BMvLjK7YJRMv8_zc\"","id":{"kind":"youtube#video","videoId":"y882AFjrSOM"},"snippet":{"publishedAt":"2016-07-31T15:00:02.000Z","channelId":"UCoc2P5sEaBV4p173cYBzDpA","title":"HyunA(현아) - '어때? (How's this?)' Official Music Video","description":"HyunA(현아) - '어때? (How's this?)' Official Music Video.","thumbnails":{"default":{"url":"https://i.ytimg.com/vi/y882AFjrSOM/default.jpg","width":120,"height":90},"medium":{"url":"https://i.ytimg.com/vi/y882AFjrSOM/mqdefault.jpg","width":320,"height":180},"high":{"url":"https://i.ytimg.com/vi/y882AFjrSOM/hqdefault.jpg","width":480,"height":360}},"channelTitle":"HyunA 현아 (Official YouTube Channel)","liveBroadcastContent":"none"}}]}`))
	testTerm := "test"
	testVideos := youtube.QueryYT(&http.Client{}, testTerm)

	s.Equal(1, len(testVideos["items"].([]interface{})))
}

func (s *YoutubeTestSuite) TestCreateVideosReturnsVideos() {
	testBuffer := bytes.NewBufferString(`{"kind":"youtube#searchListResponse","etag":"\"I_8xdZu766_FSaexEaDXTIfEWc0/SOzT2IaFjcn5yxsN-DdBk41zd4E\"","nextPageToken":"CAEQAA","regionCode":"US","pageInfo":{"totalResults":1000000,"resultsPerPage":1},"items":[{"kind":"youtube#searchResult","etag":"\"I_8xdZu766_FSaexEaDXTIfEWc0/GEeQVBHgE63BMvLjK7YJRMv8_zc\"","id":{"kind":"youtube#video","videoId":"y882AFjrSOM"},"snippet":{"publishedAt":"2016-07-31T15:00:02.000Z","channelId":"UCoc2P5sEaBV4p173cYBzDpA","title":"HyunA(현아) - '어때? (How's this?)' Official Music Video","description":"HyunA(현아) - '어때? (How's this?)' Official Music Video.","thumbnails":{"default":{"url":"https://i.ytimg.com/vi/y882AFjrSOM/default.jpg","width":120,"height":90},"medium":{"url":"https://i.ytimg.com/vi/y882AFjrSOM/mqdefault.jpg","width":320,"height":180},"high":{"url":"https://i.ytimg.com/vi/y882AFjrSOM/hqdefault.jpg","width":480,"height":360}},"channelTitle":"HyunA 현아 (Official YouTube Channel)","liveBroadcastContent":"none"}}]}`)
	testResults := youtube.Results{}
	_ = json.NewDecoder(testBuffer).Decode(&testResults)
	testVideos := youtube.CreateVideos(testResults)
	s.Equal(1, len(testVideos))
}

func TestYoutubeTestSuite(t *testing.T) {
	suite.Run(t, new(YoutubeTestSuite))
}
