package feeds

import (
	"encoding/xml"
	"net/http"
)

// RSS struct is used to expose the structure of RSS and map it into struct.
type RSS struct {
	XMLName    xml.Name `xml:"rss"`
	Text       string   `xml:",chardata"`
	Version    string   `xml:"version,attr"`
	Dc         string   `xml:"dc,attr"`
	Atom       string   `xml:"atom,attr"`
	Sy         string   `xml:"sy,attr"`
	Content    string   `xml:"content,attr"`
	Googleplay string   `xml:"googleplay,attr"`
	Itunes     string   `xml:"itunes,attr"`
	Channel    struct {
		Text string `xml:",chardata"`
		Link []struct {
			Text  string `xml:",chardata"`
			Rel   string `xml:"rel,attr"`
			Type  string `xml:"type,attr"`
			Href  string `xml:"href,attr"`
			Title string `xml:"title,attr"`
		} `xml:"link"`
		Title         string `xml:"title"`
		Generator     string `xml:"generator"`
		NewFeedURL    string `xml:"new-feed-url"`
		Description   string `xml:"description"`
		Copyright     string `xml:"copyright"`
		Language      string `xml:"language"`
		PubDate       string `xml:"pubDate"`
		LastBuildDate string `xml:"lastBuildDate"`
		Image         struct {
			Text  string `xml:",chardata"`
			Href  string `xml:"href,attr"`
			URL   string `xml:"url"`
			Title string `xml:"title"`
			Link  string `xml:"link"`
		} `xml:"image"`
		Category struct {
			Text     string `xml:",chardata"`
			AttrText string `xml:"text,attr"`
		} `xml:"category"`
		Author   string `xml:"author"`
		Summary  string `xml:"summary"`
		Explicit string `xml:"explicit"`
		Block    string `xml:"block"`
		Type     string `xml:"type"`
		Subtitle string `xml:"subtitle"`
		Keywords string `xml:"keywords"`
		Owner    struct {
			Text  string `xml:",chardata"`
			Name  string `xml:"name"`
			Email string `xml:"email"`
		} `xml:"owner"`
		Complete string `xml:"complete"`
		Item     []struct {
			Text        string `xml:",chardata"`
			Title       string `xml:"title"`
			Episode     string `xml:"episode"`
			EpisodeType string `xml:"episodeType"`
			Block       string `xml:"block"`
			Guid        struct {
				Text        string `xml:",chardata"`
				IsPermaLink string `xml:"isPermaLink,attr"`
			} `xml:"guid"`
			Link        string `xml:"link"`
			Description string `xml:"description"`
			Encoded     string `xml:"encoded"`
			PubDate     string `xml:"pubDate"`
			Author      string `xml:"author"`
			Enclosure   struct {
				Text   string `xml:",chardata"`
				URL    string `xml:"url,attr"`
				Length string `xml:"length,attr"`
				Type   string `xml:"type,attr"`
			} `xml:"enclosure"`
			Image struct {
				Text string `xml:",chardata"`
				Href string `xml:"href,attr"`
			} `xml:"image"`
			Duration string `xml:"duration"`
			Summary  string `xml:"summary"`
			Subtitle string `xml:"subtitle"`
			Keywords string `xml:"keywords"`
			Explicit string `xml:"explicit"`
		} `xml:"item"`
	} `xml:"channel"`
}

// GetFeed is used to get the feed xml by url.
func GetFeed(feedURL string) (*RSS, error) {
	res, err := http.Get(feedURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var rss *RSS
	err = xml.NewDecoder(res.Body).Decode(&rss)
	return rss, err
}
