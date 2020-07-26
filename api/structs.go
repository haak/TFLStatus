package api

import (
	"encoding/xml"
	"time"
)

// Config represents user supplied config information
type Config struct {
	AppKey  string `json:"App_Key"`
	
}

type line struct {
}

// Mode represents a mode of transport on the api
type Mode struct {
	TflType            string `json:"$type"`
	IsTflService       bool   `json:"isTflService"`
	IsFarePaying       bool   `json:"isFarePaying"`
	IsScheduledService bool   `json:"isScheduledService"`
	ModeName           string `json:"modeName"`
}

// Modes represents a list of Modes
type Modes []Mode

// Line represents a  Tube Line
type Line struct {
	Name string
	ID   string
}

// Lines are a list of Tube Lines
type Lines []Line

// PredictionDetailed represents a response from prediciton detailed api url
// type PredictionDetailed struct {
// 	Disclaimer  string
// 	WhenCreated string
// 	Line        string
// 	LineName    string
// 	Station     Station
// }

// ########################################################################

// PredictionSummary is a response to the PredictionSummary request
type PredictionSummary struct {
	XMLName  xml.Name         `xml:"ROOT"`
	Text     string           `xml:",chardata"`
	Time     Time             `xml:"Time"`
	Stations []SummaryStation `xml:"S"`
}

// Time represents a thing
type Time struct {
	Text      string `xml:",chardata"`
	TimeStamp string `xml:"TimeStamp,attr"`
}

// PredictionDetailed does a thing
type PredictionDetailed struct {
	XMLName     xml.Name          `xml:"ROOT"`
	Text        string            `xml:",chardata"`
	Xsd         string            `xml:"xsd,attr"`
	Xsi         string            `xml:"xsi,attr"`
	Xmlns       string            `xml:"xmlns,attr"`
	Disclaimer  string            `xml:"Disclaimer"`
	WhenCreated string            `xml:"WhenCreated"`
	Line        string            `xml:"Line"`
	LineName    string            `xml:"LineName"`
	S           []DetailedStation `xml:"S"`
}

// Station does a thing
type DetailedStation struct {
	Text        string             `xml:",chardata"`    //
	Code        string             `xml:"Code,attr"`    // A code representing the station (see Appendix B for valid values)
	Mess        string             `xml:"Mess,attr"`    //
	Name        string             `xml:"N,attr"`       // The Name of the station
	CurrentTime string             `xml:"CurTime,attr"` //The time the service was run in the format HH:MM:SS
	P           []DetailedPlatform `xml:"P"`
}

// Platform does a thing
type DetailedPlatform struct {
	Text      string          `xml:",chardata"`
	Name      string          `xml:"N,attr"`
	Num       string          `xml:"Num,attr"`       //The platform number
	TrackCode string          `xml:"TrackCode,attr"` //The track code of the section of track at the front of the platform
	NextTrain string          `xml:"NextTrain,attr"`
	T         []DetailedTrain `xml:"T"`
}

// Train does things
type DetailedTrain struct {
	Text           string `xml:",chardata"`           //
	SetNo          string `xml:"SetNo,attr"`          //
	TripNo         string `xml:"TripNo,attr"`         //
	SecondsTo      string `xml:"SecondsTo,attr"`      //A value representing the ‘time to station’ for this train in seconds in the format SSS
	TimeTo         string `xml:"TimeTo,attr"`         //A value representing the ‘time to station’ for this train in minutes and seconds in the format MM:SS
	Location       string `xml:"Location,attr"`       //The current location of the train
	Destination    string `xml:"Destination,attr"`    //The name of the destination of the train
	DestCode       string `xml:"DestCode,attr"`       //A code representing the destination of the train
	Order          string `xml:"Order,attr"`          //Not Assigned. Value will default to zero
	DepartTime     string `xml:"DepartTime,attr"`     //Time train departed the platform
	DepartInterval string `xml:"DepartInterval,attr"` //Interval in seconds between the departure of the specified train and the previous train
	Departed       string `xml:"Departed,attr"`       //Boolean value to determine if the train has departed the platform
	Direction      string `xml:"Direction,attr"`      //Direction of Travel
	IsStalled      string `xml:"IsStalled,attr"`      //Not Assigned. Value will default to zero
	TrainID        string `xml:"TrainId,attr"`        //
	LCID           string `xml:"LCID,attr"`           //
	InputDest      string `xml:"InputDest,attr"`      //
	TrackCode      string `xml:"TrackCode,attr"`      //The current section of track the train occupies
	LN             string `xml:"LN,attr"`             //A code representing the line the train is running on (see Appendix A)
	LeadingCarNo   string `xml:"LeadingCarNo,attr"`
}

type SummaryStation struct {
	Text string            `xml:",chardata"`
	Code string            `xml:"Code,attr"`
	N    string            `xml:"N,attr"`
	P    []SummaryPlatform `xml:"P"`
}

type SummaryPlatform struct {
	Text string          `xml:",chardata"`
	N    string          `xml:"N,attr"`
	Code string          `xml:"Code,attr"`
	Next string          `xml:"Next,attr"`
	T    []SummmaryTrain `xml:"T"`
}

type SummmaryTrain struct {
	Text string `xml:",chardata"`
	S    string `xml:"S,attr"`
	T    string `xml:"T,attr"`
	D    string `xml:"D,attr"`
	C    string `xml:"C,attr"`
	L    string `xml:"L,attr"`
	DE   string `xml:"DE,attr"`
}

// https://api.tfl.gov.uk/StopPoint/490008376N/arrivals

type AutoGenerated []struct {
	Type                string    `json:"$type"`
	ID                  string    `json:"id"`
	OperationType       int       `json:"operationType"`
	VehicleID           string    `json:"vehicleId"`
	NaptanID            string    `json:"naptanId"`
	StationName         string    `json:"stationName"`
	LineID              string    `json:"lineId"`
	LineName            string    `json:"lineName"`
	PlatformName        string    `json:"platformName"`
	Direction           string    `json:"direction"`
	Bearing             string    `json:"bearing"`
	DestinationNaptanID string    `json:"destinationNaptanId"`
	DestinationName     string    `json:"destinationName"`
	Timestamp           time.Time `json:"timestamp"`
	TimeToStation       int       `json:"timeToStation"`
	CurrentLocation     string    `json:"currentLocation"`
	Towards             string    `json:"towards"`
	ExpectedArrival     time.Time `json:"expectedArrival"`
	TimeToLive          time.Time `json:"timeToLive"`
	ModeName            string    `json:"modeName"`
	Timing              Timing    `json:"timing"`
}
type Timing struct {
	Type                      string    `json:"$type"`
	CountdownServerAdjustment string    `json:"countdownServerAdjustment"`
	Source                    time.Time `json:"source"`
	Insert                    time.Time `json:"insert"`
	Read                      time.Time `json:"read"`
	Sent                      time.Time `json:"sent"`
	Received                  time.Time `json:"received"`
}
