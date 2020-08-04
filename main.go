package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/awesome-gocui/gocui"
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

	// listModes := *boolPtr

	flag.Parse()
	// log.Info(*wordPtr)
	// log.Info("\n")
	// log.Info(*modesBoolPtr)
	// api.ListModes()
	// station := *wordPtr
	// log.Info("station: ", station)
	// api.GetSummaryPrediction(station)

	// marshallMap(stationCodeToName, "stations.json")

	// api.SaveJSON("stations.json", stationCodeToName)
	// api.SaveJSON("lines.json", lineCodeToName)

	// CloseApp()

	// Calling -modes
	if *modesBoolPtr == true {
		//  call func to list modes
		client.ModeService.GetModes()
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

	gui()

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

func gui() {
	g, err := gocui.NewGui(gocui.OutputNormal, false)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true
	g.Mouse = true

	// This sets the manager func and deletes all keybindings and views
	g.SetManagerFunc(layout)

	// Call keydbindings to set keybindings for views
	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && !gocui.IsQuit(err) {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("modes", 0, 0, maxX/4, maxY/2, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = "Modes"
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "tube")

		if _, err := g.SetCurrentView("modes"); err != nil {
			return err
		}

	}

	if v, err := g.SetView("lines", 0, maxY/2, maxX/4, maxY, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = "Lines"
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "tube")

		if _, err := g.SetCurrentView("modes"); err != nil {
			return err
		}

	}



	if v, err := g.SetView("main", maxX/4, 0, maxX*3/4, maxY, 0 ); err != nil{
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = "Main"
		// v.Highlight = true
		// v.SelBgColor = gocui.ColorGreen
		// v.SelFgColor = gocui.ColorBlack
		// fmt.Fprintln(v, "tube")
		log.Info("test")

	}

	if v, err := g.SetView("hello", maxX/2-10, maxY/2, maxX/2+10, maxY/2+5, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		fmt.Fprintln(v, "Hello world!")
		if _, err := g.SetCurrentView("hello"); err != nil {
			return err
		}
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// Plan for GUI
// have a tab on the left for lines
// then show stations somewhere
// then show trains for that station

func keybindings(g *gocui.Gui) error {

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	return nil
}
