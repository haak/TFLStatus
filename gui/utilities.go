package gui

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
	"github.com/haak/tflMap/api"
)

var (
	// viewArr = []string{"modes", "lines", "main", "hello"}
	viewArr         = []string{"modes", "lines", "main", "unknown"}
	activeViewIndex = 0
)

func nextView(g *gocui.Gui, v *gocui.View) error {
	nextViewIndex := (activeViewIndex + 1) % len(viewArr)
	name := viewArr[nextViewIndex]
	_, err := g.View("main")
	if err != nil {
		return err
	}
	// fmt.Fprintln(out, "Going from view "+v.Name()+" to "+name)

	if _, err := setCurrentViewOnTop(g, name); err != nil {
		return err
	}

	if nextViewIndex == 0 || nextViewIndex == 3 {
		g.Cursor = true
	} else {
		g.Cursor = false
	}

	activeViewIndex = nextViewIndex
	return nil
}

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func WriteModes(v *gocui.View) {

	// get gui
	// get modes view
	// get modes data
	// write to screen
	client := api.NewClient(nil)
	modes := client.ModeService.GetModesAlone()
	v.Clear()

	for _, mode := range *modes {
		fmt.Fprintln(v, mode.ModeName)

	}

}

func writeLines(v *gocui.View, mode string) {
	client := api.NewClient(nil)
	modeLines := client.LineService.GetLinesForMode(mode)
	v.Clear()

	for _, line := range *modeLines{
		fmt.Fprintln(v, line.ID)
	}


	// create client
	// get lines for mode
	//

	// for _ , line := range *lines {
	// fmt.Fprintln(v, line.name)
	// }
}

// func getLine(g *gocui.Gui, v *gocui.View) error {
// 	var l string
// 	var err error

// 	_, cy := v.Cursor()
// 	if l, err = v.Line(cy); err != nil {
// 		l = ""
// 	}

// 	maxX, maxY := g.Size()
// 	if v, err := g.SetView("msg", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2, 0); err != nil {
// 		if !gocui.IsUnknownView(err) {
// 			return err
// 		}
// 		fmt.Fprintln(v, l)
// 		if _, err := g.SetCurrentView("msg"); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

func getLine(g *gocui.Gui, v *gocui.View) (string, error) {
	_, cy := v.Cursor()
	if l, err := v.Line(cy); err != nil {
		l = ""
	} else {
		return l, err
	}
	// return l
	return "", nil

}

func pressEnter(g *gocui.Gui, v *gocui.View) error {
	// name :=
	// log.Info("test")
	if v.Name() == "modes" {

		// find the name of the mode and pass to getLine
		// getLine(g, v)
		// log.Info(l, err)

		line, _ := getLine(g, v)
		// log.Info(line)
		// out, _ := g.View("unknown")
		// fmt.Fprintln(out, "Line: "+line)
		if line == "tube"{
			linesView, _ := g.View("lines")
			// g.SetCurrentView()
			writeLines(linesView, line)
		}
	}

	if v.Name() == "lines"{
		line, _ := getLine(g, v)
		arrivalsView , _ := g.View("main")
		getArrivals(arrivalsView, line)

		// line, _ := getLine(g, v)
		stopPointView , _ := g.View("unknown")
		getStopPoints(stopPointView, line)

	}
	return nil

}


func getArrivals(v *gocui.View , line string ){
	v.Clear()
	client := api.NewClient(nil)
	fmt.Fprintln(v, line )
	lineArrivals, _, _ := client.LineService.LineArrivals(line)
	fmt.Fprintln(v, lineArrivals )
	// for _, arrival := range *lineArrivals{
	// 	fmt.Fprintln(v, arrival )
	// }

}


func getStopPoints(v *gocui.View , line string){
	v.Clear()
	client := api.NewClient(nil)
	fmt.Fprintln(v, line )
	lineStopPoints  := client.LineService.GetStopPoints(line)
	fmt.Fprintln(v, lineStopPoints )

}

