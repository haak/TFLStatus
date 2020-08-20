// For this im going to use gocui for the userinterface and the
// gocui
// log logrus
// sling "github.com/dghubble/sling"
// im not going to need anything else
// i should be able to create the map using the gocui thing
// i think ill make like a wireframe map that can be used to show all the trains or just the stations and trains nearby

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/haak/tflMap/api"
	"github.com/haak/tflMap/gui"
	log "github.com/sirupsen/logrus"
	// "github.com/jroimartin/gocui"
)

var (
	// viewArr = []string{"modes", "lines", "main", "hello"}
	viewArr         = []string{"modes", "lines", "main", "unknown"}
	activeViewIndex = 0
)

func main() {

	
	fonts()
	// fmt.Println("✓ Hello, 世界")
    // // OR
    // fmt.Println("\u2713 Hello, 世界")

	stationCodeToName := make(map[string]string)
	stationCodeToName["BST"] = "Baker Street"

	lineCodeToName := make(map[string]string)
	lineCodeToName["V"] = "Victoria"

	client := api.NewClient(nil)
	client.ModeService.GetModesAlone()
	// client.LineService.GetLineInformation("W")
	// api.GetLineInfo()
	// client.LineService.LineArrivals("victoria")
	// client.LineService.GetStopPoints("victoria")
	// client.LineService.TimetableForStop("victoria","940GZZLUBLR" )
	test := true

	if test == false {

	}

	// wordPtr := flag.String("line", "H", "London Tube line")
	modesBoolPtr := flag.Bool("modes", false, "if true list available modes")

	linesBoolPtr := flag.Bool("lines", false, "if true list available lines")

	modesStringPtr := flag.String("mode", "", "the mode to display")

	linesStringPtr := flag.String("line", "", "the line to display")

	stationsBoolPtr := flag.Bool("stations", false, "if true list stations")

	debugBoolPtr := flag.Bool("debug", false, "this is used to enable debug features")

	guiBoolPtr := flag.Bool("gui", false, "this is used to enable gui features")

	logoBoolPtr := flag.Bool("logo", false, "this is used to enable logo features")

	// listModes := *boolPtr

	flag.Parse()
	// log.Info(*wordPtr)
	// log.Info("\n")
	// log.Info(*modesBoolPtr)
	// api.ListModes()
	// station := *wordPtr2
	// log.Info("station: ", station)
	// api.GetSummaryPrediction(station)

	// marshallMap(stationCodeToName, "stations.json")

	// api.SaveJSON("stations.json", stationCodeToName)
	// api.SaveJSON("lines.json", lineCodeToName)

	// CloseApp()

	// Calling -modes
	if *modesBoolPtr == true {
		//  call func to list modes
		client.ModeService.GetModesAlone()
		// log.Info("help")
		// log.Info("mode string pointer: ", *modesStringPtr)

	}

	// Calling -mode with a mode of transport available from -modes
	if *modesStringPtr != "" {
		log.Info(*modesStringPtr)
		client.LineService.GetLinesForMode(*modesStringPtr)
	}

	// Calling -lines
	if *linesBoolPtr == true {
		// call func to list lines
		log.Info("line string pointer: ", *linesStringPtr)
	}

	if *stationsBoolPtr == true {
		// Get stations for line
	}

	if *debugBoolPtr == true {
		// do debug things here
	}

	if *guiBoolPtr == true {
		gui.CreateGui()
	}

	if *logoBoolPtr == true {
		printLogo()
	}

}

// when looking for a train line
// check storage
// if timestamp is older than 30 seconds send a new request and save in storage

// CloseApp does a thing
func CloseApp() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}

//https://api.digital.tfl.gov.uk/Line/victoria

// func marshallMap(m  map[string]string, path string) error{
// 	data, err := json.Marshal(m)

// 	f, err := os.OpenFile("json/"+path, os.O_WRONLY|os.O_CREATE, 0600)
// 	if err != nil {
// 		log.Error("error saving", path, err)
// 		return err
// 	}
// 	if err = json.NewEncoder(f).Encode(data); err != nil {
// 		log.Error("error saving", path, err)
// 		return err
// 	}
// 	return nil

// }

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

// https://api.tfl.gov.uk/Line/Route?ids=Bakerloo&serviceTypes=Regular
// dont know exactly what this returns

