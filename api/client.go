package api

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dghubble/sling"
	log "github.com/sirupsen/logrus"
)


var (
	conf = new(Config)
)

const baseURL = "https://api.tfl.gov.uk"
const baseURL2 = "http://cloud.tfl.gov.uk/TrackerNet"

// Client is a TFL api client
type Client struct {
	apiKey        string
	applicationID string
	ModeService   *ModeService
	LineService   *LineService
	// Other service end points go here
}

func (c Client) getModes(m Modes) {
	// https://api.tfl.gov.uk/Line/Meta/Modes

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

// GetModesAlone returns the modes from TFL api
func (m *ModeService) GetModesAlone() {
	modes := m.GetModes()
	for mode, element := range *modes {
		log.Info(element, " : ", mode)
		// log.Info(element)

	}

}

// ###############################################

// LineService represents a Line
type LineService struct {
	sling *sling.Sling
}

// NewLineService represents things
func NewLineService(httpClient *http.Client) *LineService {
	return &LineService{
		sling: sling.New().Client(httpClient).Base(baseURL2),
	}
}

// GetLineInformation
// func (l *LineService)GetLineInformation(line string){
// 	path := "PredictionSummary/" + line
// 	PredictionSummary := new(PredictionSummary)
// 	error := new(TFLError)
// resp, err := l.sling.New().Path(path).Receive()
// 	log.Info(resp)
// 	log.Info(err)

// }

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


func getRequest(URL string ) (http.Response, error) {
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
func GetSummaryPrediction(line string  ) (prediction PredictionSummary) {
	summaryURL := "/PredictionSummary"
	url := baseURL2 + summaryURL + "/" + line
	getRequest(url)
	log.Info(url)
	// resp, err := http.Get(url)
	// log.Info(resp)
	// log.Info(err)
	return prediction
}

// GetLines will return a string which contains all the tube lines and their code
func GetLines() string{
	line := ""
	code := ""
	// This will return a list of train lines and their codes
	log.Info("Line: " + line + " Code: " + code)
	return line
}



// https://api.tfl.gov.uk/StopPoint/490008376N/arrivals

