package easyconfig // fork appconf

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const CONFIG_FILE = "config.json"

// Parse - Parses configuration file.
func Parse(path string, config interface{}) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(config)
	return
}

// Save - Saves configuration file.
func Save(path string, config interface{}) (err error) {
	err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return
	}
	file, err := os.Create(path)
	if err != nil {
		return
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(config)
	return
}
