package api

import (
	"encoding/json"
	"os"

	// "encoding/json"
	// "os"
	log "github.com/sirupsen/logrus"
)

func loadPredictionSummary() error {
	PredictionSummary := PredictionSummary{}
	return LoadJSON("config.json", PredictionSummary)
}

func savePredictionSummary() error {
	PredictionSummary := PredictionSummary{}
	return SaveJSON("config.json", PredictionSummary)
}

// func cleanup() {
// 	for _, f := range []func() error{saveConfig, saveRelationships} {
// 		if err := f(); err != nil {
// 			log.Error("error cleaning up files", err)
// 		}
// 	}
// 	log.Info("Done cleanup. Exiting.")
// }

func LoadJSON(path string, v interface{}) error {
	f, err := os.OpenFile("json/"+path, os.O_RDONLY, 0600)
	if err != nil {
		log.Error("error loading", path, err)
		return err
	}

	if err := json.NewDecoder(f).Decode(v); err != nil {
		log.Error("error loading", path, err)
		return err
	}
	return nil
}

func SaveJSON(path string, data interface{}) error {
	f, err := os.OpenFile("json/"+path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Error("error saving", path, err)
		return err
	}

	if err = json.NewEncoder(f).Encode(data); err != nil {
		log.Error("error saving", path, err)
		return err
	}
	return nil
}

func loadConfig() error {
	return LoadJSON("config.json", conf)
}
func saveConfig() error {
	return SaveJSON("config.json", conf)
}
