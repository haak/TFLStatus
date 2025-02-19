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
	_, err := m.sling.New().Path("/Line/Meta/Modes").Receive(modes, error)
	// log.Info(sling.path)
	// log.Info(resp.Body)
	if err != nil {
		log.Info(err)
	}
	// log.Info(err)
	// log.Info(modes)
	// log.Info(error)
	return modes
	// does things
}

// GetModesAlone returns the modes from TFL api
func (m *ModeService) GetModesAlone() *Modes {
	modes := m.GetModes()
	// for mode, element := range *modes {
		
	// 	// log.Info(element, " : ", mode)
	// 	// log.Info(element)
		
	// }
	return modes

}

// func (m *ModeService) GetModesList() {
// 	modes := m.GetModes()

// 	for mode, element := range *modes {
// 		log.Info(element, " : ", mode)
// 		// log.Info(element)

// 	}
// }

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

func (m *Modes) String() string {
	output := ""
	output += "modes available \n"
	for _, mode := range *m {
		output += "Name: " + mode.ModeName + "\n"
		// output += "\n"
	}
	return output
}
