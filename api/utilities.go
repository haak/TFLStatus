package api

import (
	"strconv"
)

func (a *ArrivalsResponse) String() string {

	// This maybe needs to only show one vehicle id 


	// colorReset := "\033[0m"

	// colorRed := "\033[31m"
	// colorGreen := "\033[32m"
	// colorYellow := "\033[33m"
	// colorBlue := "\033[34m"
	// colorPurple := "\033[35m"
	// colorCyan := "\033[36m"
	// colorWhite := "\033[37m"
	// fmt.Println(string(colorRed), "test")
	// fmt.Println(string(colorReset))
	string := ""
	for _, arrivals := range *a {
		string += "stationName: " + arrivals.StationName + "\n"
		string += "vehicleId: " + arrivals.VehicleID + "\n"
		string += "naptanID: " + arrivals.NaptanID + "\n"
		string += "lineID: " + arrivals.LineID + "\n"
		string += "destination: " + arrivals.DestinationName + "\n"
		string += "Time To Station: " + strconv.Itoa(arrivals.TimeToStation) + "\n"
		string += "direction: " + arrivals.Direction + "\n"
		string += "\n"

	}
	return string
}

// JSON returns a string of json
func (a *ArrivalsResponse) JSON() string {
	return ""
}

func (l *LineStopPoints) String() string {
	output := ""
	output += "Stop Points for a Line \n"
	for _, stopPoint := range *l {
		output += "Name: " + stopPoint.CommonName + "\n"
		output += "NaptanID: " + stopPoint.NaptanID + "\n"
		// string += "NaptanID: " + stopPoint.Lines.s + "\n"
		// need to add in option for multiple lines that appear at a station

		output += "Latitude: " + strconv.FormatFloat(stopPoint.Lat, 'f', -1, 64) + "\n"
		output += "Longitude: " + strconv.FormatFloat(stopPoint.Lon, 'f', -1, 64) + "\n"
		output += "\n"

	}
	return output
}
