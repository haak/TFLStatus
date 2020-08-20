package gui

import (
	"fmt"
	"log"

	"github.com/awesome-gocui/gocui"
)

// func CreateGui(g *gocui.Gui) {
// 	// This should be called by main in the main package and it should create the gui and return a ui with some views and keybinds set.
// 	// Ideally the keybinds would be loaded from a file so they are customizable
// 	// createKeybinds(g * gocui.Gui)

// }

func CreateGui() {

	g, err := gocui.NewGui(gocui.Output256, false)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true
	g.Mouse = true

	// This sets the manager func and deletes all keybindings and views
	g.SetManagerFunc(layout)

	// Call keydbindings to set keybindings for views
	if err := createKeybindings(g); err != nil {
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
		g.Highlight = true
		g.SelFgColor = gocui.ColorGreen

		// Text colours
		// g.FgColor = gocui.ColorBlack

		// g.SelBgColor = gocui.ColorBlack
		// This sets the fram colour

		// This changes frame colour of the selected frame
		g.SelFrameColor = gocui.ColorGreen
		// g.SelFrameColor = gocui.ColorYellow

		g.FrameColor = gocui.ColorDefault

		v.Title = "Modes"
		// v.Highlight = true
		// v.SelBgColor = gocui.ColorGreen
		// v.SelFgColor = gocui.ColorBlack

		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorDefault

		WriteModes(v)
		// fmt.Fprintln(v, "tube")

		if _, err := g.SetCurrentView("modes"); err != nil {
			return err
		}

	}

	if v, err := g.SetView("lines", 0, maxY/2+1, maxX/4, maxY-1, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = "Lines"
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		// fmt.Fprintln(v, "tube")

		if _, err := g.SetCurrentView("modes"); err != nil {
			return err
		}

	}

	if v, err := g.SetView("main", maxX/4+1, 0, maxX*3/4, maxY-1, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = "Main"
		// v.Highlight = true
		// v.SelBgColor = gocui.ColorGreen
		// v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, `
▧▧▧▧▧▧▧▧▧▧▧▧▧`)
		// log.Info("test")

	}

	if v, err := g.SetView("unknown", maxX*3/4+1, 0, maxX-1, maxY-1, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = "Unknown"
		// v.Highlight = true
		// v.SelBgColor = gocui.ColorGreen
		// v.SelFgColor = gocui.ColorBlack
		// fmt.Fprintln(v, "tube")
		// log.Info("test")

	}

	// if v, err := g.SetView("hello", maxX/2-10, maxY/2, maxX/2+10, maxY/2+5, 0); err != nil {
	// 	if !gocui.IsUnknownView(err) {
	// 		return err
	// 	}
	// 	fmt.Fprintln(v, "Hello world!")
	// 	if _, err := g.SetCurrentView("hello"); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}
