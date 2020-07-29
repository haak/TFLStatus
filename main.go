package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/haak/tflMap/api"
	log "github.com/sirupsen/logrus"
)

func main() {

	stationCodeToName := make(map[string]string)
	stationCodeToName["BST"] = "Baker Street"

	lineCodeToName := make(map[string]string)
	lineCodeToName["V"] = "Victoria"

	client := api.NewClient(nil)
	// client.ModeService.GetModesAlone()
	// client.LineService.GetLineInformation("W")
	// api.GetLineInfo()
	client.LineService.LineArrivals("victoria")
	wordPtr := flag.String("line", "H", "London Tube line")
	flag.Parse()
	log.Info(*wordPtr)
	station := *wordPtr
	log.Info("station: ", station)
	// api.GetSummaryPrediction(station)

	// marshallMap(stationCodeToName, "stations.json")

	api.SaveJSON("stations.json", stationCodeToName)
	api.SaveJSON("lines.json", lineCodeToName)

	// CloseApp()

}

// when looking for a train line
// check storage
// if timestamp is older than 30 seconds send a new request and save in storage

// CloseApp

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
