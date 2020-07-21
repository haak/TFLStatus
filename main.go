package main

import (
	"github.com/haak/tflMap/api"
)

func main() {
	client := api.NewClient(nil)
	client.ModeService.GetModesAlone()

}
