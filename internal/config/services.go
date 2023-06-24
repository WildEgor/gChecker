package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

const ConfigName = "services.json"

type ServiceUrl struct {
	ID      string `json:"id"`
	URL     string `json:"url"`
	Enabled bool   `json:"enabled"`
}

type ServicesConfig struct {
	Timeout int8         `json:"timeout"`
	URLs    []ServiceUrl `json:"urls"`
}

func NewServicesConfig() *ServicesConfig {
	var configPath string
	var config ServicesConfig

	workingdir, _ := os.Getwd()

	if len(ConfigName) == 0 {
		configPath = filepath.Join(workingdir, ConfigName)
	} else {
		configPath = filepath.Join(ConfigName)
	}

	bts, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal("[ServicesConfig] Cannot open file", err)
	}

	if err := json.Unmarshal(bts, &config); err != nil {
		log.Fatal("[ServicesConfig] Cannot unmarshal file", err)
	}

	log.Print(config)

	return &config
}
