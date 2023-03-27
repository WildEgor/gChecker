package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"time"
)

type RecordMap map[string]string

type ServiceUrl struct {
	ID        int8      `json:"id"`
	URL       string    `json:"url"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ServicesConfig struct {
	CheckTimeout int8         `json:"timeout"`
	MessTimeout  int8         `json:"messTimeout"`
	ServiceUrls  []ServiceUrl `json:"urls"`
}

const (
	ServicesConfigFileName = "services.json"
)

func NewServicesConfig() *ServicesConfig {
	var cName string
	var config ServicesConfig

	workingdir, _ := os.Getwd()

	if len(ServicesConfigFileName) == 0 {
		cName = filepath.Join(workingdir, ServicesConfigFileName)
	} else {
		cName = filepath.Join(ServicesConfigFileName)
	}

	bts, err := os.ReadFile(cName)
	if err != nil {
		log.Fatal("Cannot open file", err)
	}

	if err := json.Unmarshal(bts, &config); err != nil {
		log.Fatal("Cannot unmarshal file", err)
	}

	log.Print(config)

	return &config
}
