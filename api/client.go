package api

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var (
	conf = new(Config)
)

const baseURL = "https://api.tfl.gov.uk"

// const baseURL2 = "http://cloud.tfl.gov.uk/TrackerNet"

// Client is a TFL api client
type Client struct {
	apiKey           string
	applicationID    string
	ModeService      *ModeService
	LineService      *LineService
	StopPointService *StopPointService
	// Other service end points go here
}

// NewClient returns a new client
func NewClient(httpClient *http.Client) *Client {
	return &Client{ModeService: NewModeService(httpClient),
		LineService: NewLineService(httpClient)}
}

// TFLError represents a TFL API error response
// https://developer.github.com/v3/#client-errors
type TFLError struct {
	Message string `json:"message"`
}

// ####################################################

// GetLineInfo returns a thing (THIS USES THE OLD API)
func GetLineInfo() {
	resp, err := http.Get("http://cloud.tfl.gov.uk/TrackerNet/PredictionSummary/W")
	log.Info(&resp.Body)
	log.Info(err)
	if err != nil {
		log.Info(err)
		// print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Info(err)
		// print(err)
	}
	// fmt.Print(string(body))
	log.Info(string(body))

	PredictionSummary := PredictionSummary{}
	xml.Unmarshal(body, &PredictionSummary)
	fmt.Println("Prediction summary Time: " + PredictionSummary.Time.TimeStamp)
	for _, item := range PredictionSummary.Stations {
		fmt.Println(item.N)
	}
	for _, station := range PredictionSummary.Stations {
		fmt.Println(station.Code)
	}
	log.Info(PredictionSummary)
	// tmpBool := checkOnline(body)
	// log.Info(tmpBool)

}

// when we are making a client
// it has methods on it

func getRequest(URL string) (http.Response, error) {
	resp, err := http.Get(URL)
	log.Info(&resp.Body)
	log.Info(err)
	if err != nil {
		log.Info(err)
		// print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Info(err)
		// print(err)
	}
	// fmt.Print(string(body))
	log.Info(string(body))
	return *resp, err
}

// GetSummaryPrediction returns a prediction for a line
// func GetSummaryPrediction(line string  ) (prediction PredictionSummary) {
// 	summaryURL := "/PredictionSummary"
// 	url := baseURL2 + summaryURL + "/" + line
// 	getRequest(url)
// 	log.Info(url)
// 	// resp, err := http.Get(url)
// 	// log.Info(resp)
// 	// log.Info(err)
// 	return prediction
// }

// GetLines will return a string which contains all the tube lines and their code
func GetLines() string {
	line := ""
	code := ""
	// This will return a list of train lines and their codes
	log.Info("Line: " + line + " Code: " + code)
	return line
}

// https://api.tfl.gov.uk/StopPoint/490008376N/arrivals
