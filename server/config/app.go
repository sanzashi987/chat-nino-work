package config

import (
	// "encoding/json"
	"encoding/json"
	"os"
	"sync"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Config string `name:"config" type:"file" default:"config.json"`
}

var once sync.Once
var config *AppConfig

type AppConfig struct {
	ApiKey   string `json:"api_key"`
	Port     int    `json:"port"`
	ProxyUrl string `json:"proxy_url"`
}

func LoadConfig() *AppConfig {
	once.Do(func() {
		kong.Parse(&CLI)

		config = &AppConfig{
			ApiKey:   "",
			ProxyUrl: "",
			Port:     8080,
		}

		_, err := os.Stat(CLI.Config)
		if err != nil {
			return
		}

		file, err := os.Open(CLI.Config)
		if err != nil {
			return
		}
		defer file.Close()

		encoder := json.NewDecoder(file)
		err = encoder.Decode(config)

		if err != nil {
			return
		}

	})

	return config
}
