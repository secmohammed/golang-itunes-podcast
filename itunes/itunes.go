package itunes

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

//API struct is used to expose the itunes api.
type API struct {
}

//SearchResponse is used to show the search response
type SearchResponse struct {
	ResultCount int `json:"resultCount"`
	Results     []struct {
		WrapperType            string    `json:"wrapperType"`
		Kind                   string    `json:"kind"`
		CollectionID           int       `json:"collectionId"`
		TrackID                int       `json:"trackId"`
		ArtistName             string    `json:"artistName"`
		CollectionName         string    `json:"collectionName"`
		TrackName              string    `json:"trackName"`
		CollectionCensoredName string    `json:"collectionCensoredName"`
		TrackCensoredName      string    `json:"trackCensoredName"`
		CollectionViewURL      string    `json:"collectionViewUrl"`
		FeedURL                string    `json:"feedUrl,omitempty"`
		TrackViewURL           string    `json:"trackViewUrl"`
		ArtworkURL30           string    `json:"artworkUrl30"`
		ArtworkURL60           string    `json:"artworkUrl60"`
		ArtworkURL100          string    `json:"artworkUrl100"`
		CollectionPrice        float64   `json:"collectionPrice"`
		TrackPrice             float64   `json:"trackPrice"`
		TrackRentalPrice       int       `json:"trackRentalPrice"`
		CollectionHdPrice      int       `json:"collectionHdPrice"`
		TrackHdPrice           int       `json:"trackHdPrice"`
		TrackHdRentalPrice     int       `json:"trackHdRentalPrice"`
		ReleaseDate            time.Time `json:"releaseDate"`
		CollectionExplicitness string    `json:"collectionExplicitness"`
		TrackExplicitness      string    `json:"trackExplicitness"`
		TrackCount             int       `json:"trackCount"`
		Country                string    `json:"country"`
		Currency               string    `json:"currency"`
		PrimaryGenreName       string    `json:"primaryGenreName"`
		ContentAdvisoryRating  string    `json:"contentAdvisoryRating,omitempty"`
		ArtworkURL600          string    `json:"artworkUrl600"`
		GenreIds               []string  `json:"genreIds"`
		Genres                 []string  `json:"genres"`
		ArtistID               int       `json:"artistId,omitempty"`
		ArtistViewURL          string    `json:"artistViewUrl,omitempty"`
	} `json:"results"`
}

//NewAPI is used to construct a new api of itunes.
func NewAPI() *API {
	return &API{}
}

//Search is used to use the itunes search api to search for a term.
func (api *API) Search(term string) (SearchResponse, error) {

	searchURL := url.URL{
		Scheme: "https",
		Host:   "itunes.apple.com",
		Path:   "search",
	}
	q := searchURL.Query()
	q.Set("entity", "podcast")
	q.Set("term", term)
	searchURL.RawQuery = q.Encode()
	res, err := http.Get(searchURL.String())
	if err != nil {
		return SearchResponse{}, err
	}
	defer res.Body.Close()
	var searchResponse SearchResponse
	err = json.NewDecoder(res.Body).Decode(&searchResponse)
	return searchResponse, err
}
