package api

import (
	"net/http"

	"github.com/dghubble/sling"
	log "github.com/sirupsen/logrus"
)

const pointURL = "https://api.tfl.gov.uk/StopPoint/"

// Services

// StopPointService can be used to make requests related to the stopPoint endpoint
type StopPointService struct {
	sling *sling.Sling
}

// NewStopPointService creates new StopPointService
func NewStopPointService(httpClient *http.Client) *StopPointService {
	return &StopPointService{
		sling: sling.New().Client(httpClient).Base(pointURL),
	}

}

// PointArrivals logs the arrivals for a specific stop point
func (s StopPointService) PointArrivals(stopPoint string) (*ArrivalsResponse, *http.Response, error) {
	ArrivalsResponse := new(ArrivalsResponse)
	error := new(error)
	path := stopPoint + "/arrivals"

	resp, err := s.sling.New().Get(path).Receive(ArrivalsResponse, error)
	if err != nil {
		panic(err)
	}

	log.Info(ArrivalsResponse)
	return ArrivalsResponse, resp, err
}



