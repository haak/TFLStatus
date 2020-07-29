package api

import (
	"fmt"
	"strconv"
)

func (a *ArrivalsResponse) String() string {
	// colorReset := "\033[0m"

    colorRed := "\033[31m"
    // colorGreen := "\033[32m"
    // colorYellow := "\033[33m"
    // colorBlue := "\033[34m"
    // colorPurple := "\033[35m"
    // colorCyan := "\033[36m"
	// colorWhite := "\033[37m"
	fmt.Println(string(colorRed), "test")
	string := ""
	for _, arrivals := range *a {
		string += "vehicleId: " + arrivals.VehicleID + "\n"
		string += "naptanID: " + arrivals.NaptanID + "\n"
		string += "stationName: " + arrivals.StationName + "\n"
		string += "lineID: " + arrivals.LineID + "\n"
		string += "destination: " + arrivals.DestinationName + "\n"
		string += "Time To Station: " + strconv.Itoa(arrivals.TimeToStation) + "\n"
		string += "\n"
		
	}
	return string
}

// JSON returns a string of json
func (a *ArrivalsResponse) JSON() string {
	return ""
}
