package api

import (
	"net/http"

	"github.com/dghubble/sling"
	log "github.com/sirupsen/logrus"
)

const lineURL = "https://api.tfl.gov.uk/line/"

//https://api.tfl.gov.uk/line/victoria/arrivals

// This makes use of the ArrivalsResponse struct


//Services

type LineService struct {
	sling *sling.Sling
}

func NewLineService(httpClient *http.Client) *LineService {
	return &LineService{
		sling: sling.New().Client(httpClient).Base(lineURL),
	}
}

func (l *LineService) LineArrivals(line string) (*ArrivalsResponse, *http.Response, error) {
	ArrivalsResponse := new(ArrivalsResponse)
	error := new(error)
	path := line + "/arrivals"
	resp, err := l.sling.New().Get(path).Receive(ArrivalsResponse, error)
	if err != nil {
		panic(err)
	}
	log.Info("there seems to be no response")
	log.Info(ArrivalsResponse)
	return ArrivalsResponse, resp, err
}


func (l *LineService) getStopPoints(line string){
	//https://api.tfl.gov.uk/line/victoria/StopPoints
	///Line/{id}/StopPoints
	path := line + "/StopPoints"
	// resp, err := l.sling.New().Get(path).Receive(ArrivalsResponse, error)
	// if err != nil {
	// 	panic(err)
	// }

}