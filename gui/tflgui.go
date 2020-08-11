package gui

import (
	"fmt"

	"github.com/haak/tflMap/api"
	"github.com/awesome-gocui/gocui"
)

func WriteModes(v *gocui.View) {

	// get gui
	// get modes view
	// get modes data
	// write to screen
	client := api.NewClient(nil)
	modes := client.ModeService.GetModesAlone()

	for _, mode := range *modes{
		fmt.Fprintln(v, mode.ModeName)

	}

	
}

func ColourChange(colourA, colourB int, input string) string {
	return ""
}
