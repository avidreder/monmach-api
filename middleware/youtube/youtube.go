package youtube

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// Video is a struct representing a Youtube video resource
type Video struct {
	Title string `json:"title"`
	ID    string `json:"id"`
}

// Videos is a collection of Video resources
type Videos []Video

// Results is a struct of raw Youtube search results
type Results map[string]interface{}

// CompileVideos gets results from YT and formats them
func CompileVideos(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		client := &http.Client{}
		term := "music"
		if c.QueryParam("query") != "" {
			term = c.QueryParam("query")
		}
		rawVideos := QueryYT(client, term)
		videos := CreateVideos(rawVideos)
		c.Set("VideoList", videos)
		return h(c)
	}
}

// QueryYT searches Youtube for a search term
func QueryYT(httpClient *http.Client, searchTerm string) Results {
	baseURL := "https://www.googleapis.com/youtube/v3/search?maxResults=5&type=video&key=AIzaSyD5e2odC6GoTPF73GEF1uoN3o2OT9eCg5k&part=snippet&q="
	baseURL += searchTerm
	req, err := http.NewRequest("GET", baseURL, nil)
	req.Header.Add("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	videos := Results{}
	err = json.NewDecoder(resp.Body).Decode(&videos)
	if err != nil {
		log.Print(err)
	}
	return videos
}

// CreateVideos takes raw search results and formats them
func CreateVideos(rawVideos Results) Videos {
	videoItems := rawVideos["items"].([]interface{})
	videoSlice := Videos{}
	for _, video := range videoItems {
		videoSlice = append(videoSlice, Video{
			video.(map[string]interface{})["snippet"].(map[string]interface{})["title"].(string),
			video.(map[string]interface{})["id"].(map[string]interface{})["videoId"].(string),
		})
	}
	return videoSlice
}
