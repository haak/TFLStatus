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

func (l *LineService) GetStopPoints(line string) {
	//https://api.tfl.gov.uk/line/victoria/StopPoints
	LineStopPointsResponse := new(LineStopPoints)
	///Line/{id}/StopPoints
	error := new(error)
	path := line + "/StopPoints"
	_, err := l.sling.New().Get(path).Receive(LineStopPointsResponse, error)
	// log.Info(resp)
	log.Info(LineStopPointsResponse)
	if err != nil {
		
	}

}


func (l *LineService) TimetableForStop(stopPoint, line string){
	// /Line/{id}/Timetable/{fromStopPointId}

	// https://api.tfl.gov.uk/line/victoria/Arrivals/940GZZLUEUS
	// This uses same as 

	return 

	//Get the list of arrival predictions for given line ids based at the given stop
	// /Line/{ids}/Arrivals/{stopPointId}
}


//https://api.tfl.gov.uk/Line/victoria/Timetable/940GZZLUKSX