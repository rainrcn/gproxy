package config

import (
	"log"
)

type AppConfig struct {
	Http  HttpConfigSection  `yaml:"http"`
	Https HttpsConfigSection `yaml:"https"`
	Maps  map[string]string  `yaml:"maps"`
}

type HttpConfigSection struct {
	Enabled    bool   `yaml:"enabled"`
	ListenAddr string `yaml:"listen-addr"`
}

type HttpsConfigSection struct {
	Enabled    bool   `yaml:"enabled"`
	ListenAddr string `yaml:"listen-addr"`
	CertFile   string `yaml:"cert-file"`
	KeyFile    string `yaml:"key-file"`
}

var appConfig = AppConfig{}

func LoadAppConfig() *AppConfig {
	err := LoadConfig("./config/app.yml", &appConfig)

	if err != nil {
		log.Fatal("load 'app.yml' error:", err)
	}

	return &appConfig
}

func GetAppConfig() *AppConfig {
	return &appConfig
}
