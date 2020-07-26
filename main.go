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
	// client := api.NewClient(nil)
	// client.ModeService.GetModesAlone()
	// client.LineService.GetLineInformation("W")
	// api.GetLineInfo()
	wordPtr := flag.String("line", "H", "London Tube line")
	flag.Parse()
	log.Info(*wordPtr)
	station := *wordPtr
	api.GetSummaryPrediction(station)
	CloseApp()
	

}

// when looking for a train line
// check storage
// if timestamp is older than 30 seconds send a new request and save in storage

// CloseApp

func CloseApp(){
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