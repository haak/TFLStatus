package api

import (
	"net/http"

	"github.com/dghubble/sling"
	log "github.com/sirupsen/logrus"
)

const baseURL = "https://api.tfl.gov.uk"
const baseURL2 = "http://example.gov.uk/TrackerNet"

// Client is a TFL api client
type Client struct {
	apiKey        string
	applicationID string
	ModeService   *ModeService
	// Other service end points go here
}

// GithubError represents a Github API error response
// https://developer.github.com/v3/#client-errors
type TFLError struct {
	Message string `json:"message"`
}

func (c Client) getModes(m Modes) {
	// https://api.tfl.gov.uk/Line/Meta/Modes

}

// ModeService provides methods for finding modes of transport.
type ModeService struct {
	sling *sling.Sling
}

// NewModeService returns a new ModeService
func NewModeService(httpClient *http.Client) *ModeService {
	return &ModeService{
		sling: sling.New().Client(httpClient).Base(baseURL),
	}
}

// GetModes returns a list of avaiable modes of transport
func (m *ModeService) GetModes() *Modes {
	// returns []Modes, *http.Response, error
	modes := new(Modes)
	error := new(TFLError)
	resp, err := m.sling.New().Path("/Line/Meta/Modes").Receive(modes, error)
	log.Info(resp.Body)
	log.Info(err)
	log.Info(modes)
	log.Info(error)
	return modes
	// does things
}

// NewClient returns a new client
func NewClient(httpClient *http.Client) *Client {
	return &Client{ModeService: NewModeService(httpClient)}
}

func (m *ModeService) GetModesAlone() {
	modes := m.GetModes()
	for mode, element := range *modes {
		log.Info(element, " : ", mode)
		// log.Info(element)

	}

}

///Mode/{mode}/Arrivals

// take a station name
// show timetable for that station


//This section is going to represent the TracketNet API that is available at 
// 4 urls
// prediction service summary
// prediction services detailed
// station status (same as line status usage)
// line status (this could be used to change text colour of station maybe red if its behind or something)






type LineService struct {
	sling *sling.Sling
}

func NewLineService(httpClient *http.Client) * LineService{
	return &LineService{
		sling: sling.New().Client(httpClient).Base(baseURL2),
	}
}
