package api

import (
	"net/http"

	"github.com/dghubble/sling"
	log "github.com/sirupsen/logrus"
)

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

// NewLineService represents things
// func NewLineService(httpClient *http.Client) *LineService {
// 	return &LineService{
// 		sling: sling.New().Client(httpClient).Base(baseURL2),
// 	}
// }

// GetLineInformation
// func (l *LineService)GetLineInformation(line string){
// 	path := "PredictionSummary/" + line
// 	PredictionSummary := new(PredictionSummary)
// 	error := new(TFLError)
// resp, err := l.sling.New().Path(path).Receive()
// 	log.Info(resp)
// 	log.Info(err)

// }
