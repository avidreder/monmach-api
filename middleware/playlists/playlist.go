package bitshows

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/labstack/echo"
	"golang.org/x/net/html"
)

// BITShow represents all of the data for a BIT show
type BITShow struct {
	DateTime  string      `json:"dateTime,omitempty"`
	Date      string      `json:"date,omitempty"`
	Time      string      `json:"time,omitempty"`
	ID        int         `json:"id,omitempty"`
	TicketURL string      `json:"ticket_url,omitempty"`
	Artists   []BITArtist `json:"artists,omitempty"`
	Venue     BITVenue    `json:"venue,omitempty"`
	VenueName string      `json:"venueName,omitempty"`
	Address   string      `json:"address,omitempty"`
	ImageURL  string      `json:"image_url,omitempty"`
}

// Show represents all of the data for a Showhawk show
type Show struct {
	Date         string   `json:"date"`
	Time         string   `json:"time"`
	ID           int      `json:"id"`
	TicketURL    string   `json:"ticket_url"`
	Artists      []string `json:"bands"`
	Venue        string   `json:"venue"`
	Address      string   `json:"address"`
	ImageURL     string   `json:"image"`
	Popularity   int      `json:"popularity"`
	Neighborhood string   `json:"neighborhood"`
	Price        int      `json:"price"`
}

// Shows are a collection of Show structs
type Shows []Show

// BITArtist represents a BIT Artist's data
type BITArtist struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
}

// BITVenue represents a BIT Artist's data
type BITVenue struct {
	ID        int     `json:"id"`
	URL       string  `json:"url"`
	Name      string  `json:"name"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Address   string  `json:"address"`
}

// SearchParameters are a stuct for all parameters used in a BIT Query
type SearchParameters struct {
	Radius   int    `url:"radius"`
	PerPage  int    `url:"per_page"`
	Date     string `url:"date"`
	Format   string `url:"format"`
	AppID    string `url:"app_id"`
	Location string `url:"location"`
}

// BITShows is a collection of BITShow structs
type BITShows []BITShow

// CompileShows prepares the Query to BIT and returns the results
func CompileShows(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		date := time.Now().String()
		dateString := date[:10]
		client := &http.Client{}
		if c.QueryParam("date") != "" {
			dateString = c.QueryParam("date")
		}
		rawShows := QueryBIT(client, dateString)
		rawShows = ParseShowDates(rawShows)
		rawShows = GetShowImages(client, rawShows)
		rawShows = GetShowAddresses(client, rawShows)
		shows, err := MapBITShowsToShows(rawShows)
		if err != nil {
			log.Print(err)
		}
		c.Set("ShowList", shows)
		return h(c)
	}

}

// MapBITShowsToShows transforms BITShows to Shows
func MapBITShowsToShows(bShows BITShows) (Shows, error) {
	shShows := make(Shows, 0)
	for i := 0; i < len(bShows); i++ {
		var artists []string
		for j := 0; j < len(bShows[i].Artists); j++ {
			artists = append(artists, bShows[i].Artists[j].Name)
		}
		shShows = append(shShows, Show{
			bShows[i].Date,
			bShows[i].Time,
			bShows[i].ID,
			bShows[i].TicketURL,
			artists,
			bShows[i].VenueName,
			bShows[i].Address,
			bShows[i].ImageURL,
			0,
			"",
			0,
		})
	}
	return shShows, nil
}

// QueryBIT searches BIT for shows
func QueryBIT(httpClient *http.Client, showDate string) BITShows {

	params := SearchParameters{
		10, 10, showDate, "json", "showhawk", "Portland,OR",
	}
	v, _ := query.Values(params)

	baseURL := "http://api.bandsintown.com/events/search?"
	baseURL += v.Encode()
	req, err := http.NewRequest("GET", baseURL, nil)
	req.Header.Add("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	shows := make(BITShows, 0)
	// log.Print(ioutil.ReadAll(resp.Body))
	err = json.NewDecoder(resp.Body).Decode(&shows)
	if err != nil {
		log.Print(err)
	}
	return shows
}

// ParseShowDates takes a show's dateTime and separates the two components
func ParseShowDates(shows BITShows) BITShows {
	for i := 0; i < len(shows); i++ {
		dateString, timeString, err := ParseDateAndTime(shows[i].DateTime)
		if err != nil {
			log.Print(err)
		}
		shows[i].Date = dateString
		shows[i].Time = timeString
	}
	return shows
}

// GetShowImages gets address and images for search results
func GetShowImages(httpClient *http.Client, shows BITShows) BITShows {
	for i := 0; i < len(shows); i++ {
		shows[i].ImageURL = "https://pixabay.com/static/uploads/photo/2015/06/08/15/24/hawk-802054_960_720.jpg"
		for j := 0; j < len(shows[i].Artists); j++ {
			shows[i].Artists[j].ImageURL = GetImageURL(httpClient, shows[i].Artists[j].URL)
			if shows[i].Artists[j].ImageURL != "https://s3.amazonaws.com/bit-photos/artistLarge.jpg" {
				shows[i].ImageURL = shows[i].Artists[j].ImageURL
			}
		}

	}
	return shows
}

// GetShowAddresses gets all the addresses of a show
func GetShowAddresses(httpClient *http.Client, shows BITShows) BITShows {
	for i := 0; i < len(shows); i++ {
		shows[i].VenueName = shows[i].Venue.Name
		shows[i].Venue.Address = GetAddress(httpClient, shows[i].Venue.Latitude, shows[i].Venue.Longitude)
		shows[i].Address = shows[i].Venue.Address
	}
	return shows
}

// ParseDateAndTime separates date and time from a string
func ParseDateAndTime(dateTimeString string) (dateString string, timeString string, err error) {
	parsedDateTime, _ := time.Parse("2006-01-02T15:04:05", dateTimeString)
	parsedDate := parsedDateTime.Format("Mon 1/2")
	parsedTime := parsedDateTime.Format("3:04 PM")
	return parsedDate, parsedTime, nil
}

// GetImageURL crawls the BIT Artist page for an image link
func GetImageURL(httpClient *http.Client, artistURL string) string {
	baseURL := artistURL
	req, err := http.NewRequest("GET", baseURL, nil)
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Print(err)
	}
	imageURL := ""
	tempContent := ""
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "meta" && imageURL == "" {
			for _, a := range n.Attr {
				if a.Key == "content" {
					tempContent = a.Val
				}
				if a.Val == "og:image" {
					imageURL = tempContent
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if imageURL != "" {
				break
			}
			f(c)
		}
	}
	f(doc)
	return imageURL
}

// GetAddress takes the coordinates of a show, and returns a street address
func GetAddress(httpClient *http.Client, latitude float32, longitude float32) string {
	addressString := ""
	baseURL := "http://maps.googleapis.com/maps/api/geocode/json?"
	baseURL += fmt.Sprintf("latlng=%v,%v", latitude, longitude)
	req, err := http.NewRequest("GET", baseURL, nil)
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	// log.Print(ioutil.ReadAll(resp.Body))
	address := map[string]interface{}{}
	// log.Print(ioutil.ReadAll(resp.Body))
	err = json.NewDecoder(resp.Body).Decode(&address)
	if err != nil {
		log.Print(err)
	}
	if len(address["results"].([]interface{})) > 0 {
		addressString = address["results"].([]interface{})[0].(map[string]interface{})["formatted_address"].(string)
	}
	return addressString
}
