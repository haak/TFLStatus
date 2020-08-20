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
	// log.Info("there seems to be no response")
	// log.Info(ArrivalsResponse)
	return ArrivalsResponse, resp, err
}

func (l *LineService) GetStopPoints(line string) *LineStopPoints {
	//https://api.tfl.gov.uk/line/victoria/StopPoints
	LineStopPointsResponse := new(LineStopPoints)
	///Line/{id}/StopPoints
	error := new(error)
	path := line + "/StopPoints"
	_, err := l.sling.New().Get(path).Receive(LineStopPointsResponse, error)

	
	// log.Info(resp)
	// log.Info(LineStopPointsResponse)
	if err != nil {
		// log.Panic("We have an error:", err)
	}
	return LineStopPointsResponse

}

// TimetableForStop returns an ArrivalResponse for the trains that are arriv
func (l *LineService) TimetableForStop(line, stopPoint string) {
	//e.g. https://api.tfl.gov.uk/line/victoria/Arrivals/940GZZLUBLR
	ArrivalsResponse := new(ArrivalsResponse)
	error := new(error)
	path := line + "/arrivals/" + stopPoint
	_, err := l.sling.New().Get(path).Receive(ArrivalsResponse, error)
	if err != nil {
		// log.Panic("We have an error:", err)
	}
	log.Info(ArrivalsResponse)
	// This uses the same thing as LineArrivals()
	// /Line/{id}/Timetable/{fromStopPointId} THIS DOES NOT WORK

	// https://api.tfl.gov.uk/line/victoria/Arrivals/940GZZLUEUS
	// WE ARE GOING TO USE THIS INSTEAD

	return

	//Get the list of arrival predictions for given line ids based at the given stop
	// /Line/{ids}/Arrivals/{stopPointId}
}

func (l *LineService) GetLinesForMode(mode string) *ModeLines{
	// /Line/Mode/{modes}
	ModeLines := new(ModeLines)
	tflError := new(TFLError)
	path := "/line/mode/" + mode
	l.sling.New().Get(path).Receive(ModeLines, tflError)
	// resp, err := l.sling.New().Get(path).Receive(ModeLines, tflError)

	// log.Info(err)
	// log.Info(resp.Body)
	// log.Info(ModeLines)
	return ModeLines

}

//https://api.tfl.gov.uk/Line/victoria/Timetable/940GZZLUKSX

func (m *ModeLines) String() string {
	output := ""
	output += "lines available \n"
	for _, modeLine := range *m {
		output += "Name: " + modeLine.Name + "\n"
		output += "ID: " + modeLine.ID + "\n"
		output += "mode: " + modeLine.ModeName + "\n"

		// output += "Route Sections: " + modeLine.RouteSections + "\n"

		// output += "Name: " + modeLine.ServiceTypes.String() + "\n"

		for _, serviceType := range modeLine.ServiceTypes {
			output += serviceType.String()
		}

		output += "\n"
	}
	return output
}

func (s ServiceTypes) String() string {
	output := ""
	output += "Service Type: " + s.Name + "\n"
	return output

}


