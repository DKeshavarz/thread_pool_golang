package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	NumWorker     int `json:"num_worker"`
	MaxQueueSize  int `json:"queue_size"`
}

func LoadConfig(filename string) (Config, error) {
	var config Config
	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return config, err
}