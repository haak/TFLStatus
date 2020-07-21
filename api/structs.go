package api

type line struct {
}

type Mode struct {
	TflType            string `json:"$type"`
	IsTflService       bool   `json:"isTflService"`
	IsFarePaying       bool   `json:"isFarePaying"`
	IsScheduledService bool   `json:"isScheduledService"`
	ModeName           string `json:"modeName"`
}

// func (m Mode) String() string {
// 	output :=
// }

type Modes []Mode

type Line struct {
	Name string
	ID   string
}

type Lines []Line

// PredictionDetailed represents a response from prediciton detailed api url
type PredictionDetailed struct {
	Disclaimer  string
	WhenCreated string
	Line        string
	LineName    string
	Station     Station
}

type Station struct {
	platforms []Platform
}

type Platform struct {
	name      string
	Number    int
	TrackCode string
	NextTrain bool
	Trips     []Trip
}

type Trip struct {
	SetNo          int
	TripNo         int
	SecondsTo      int
	TimeTo         string
	Location       string
	Destination    string
	DestCode       string
	OrderCode      string
	DepartTime     string
	DepartInterval string
	Departed       bool
	Direction      string
	IsStalled      string
}
