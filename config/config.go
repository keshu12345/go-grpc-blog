package config

import (
	"os"
	"path/filepath"

	logger "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const (
	serverYML = "server.yml"
)

type Config struct {
	EnvironmentName string `yaml:"environmentName"`
	Server          Server `yaml:"server"`
	Logger          Logger `yaml:"logger"`
}

type Server struct {
	Host         string `yaml:"Host"`
	Port         int    `yaml:"Port"`
	ReadTimeout  int    `yaml:"ReadTimeout"`
	WriteTimeout int    `yaml:"WriteTimeout"`
	IdleTimeout  int    `yaml:"IdleTimeout"`
}

type Logger struct {
	Base         string `yaml:"base"`
	Level        string `yaml:"level"`
	Filename     string `yaml:"filename"`
	MaxSize      int    `yaml:"maxsize"`
	MaxAge       int    `yaml:"maxage"`
	MaxBackups   int    `yaml:"maxbackups"`
	LocalTime    bool   `yaml:"localtime"`
	Compress     bool   `yaml:"compress"`
	ReportCaller bool   `yaml:"reportcaller"`
}

func GetConfig(path string) Config {
	var config Config
	wd, err := os.Getwd()
	if err != nil {
		logger.Fatalf("error getting working directory: %v", err)
	}
	configPath := filepath.Join(wd, path, serverYML)
	data, err := os.ReadFile(configPath)
	if err != nil {
		logger.Fatalf("error reading config file: %v", err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		logger.Fatalf("error unmarshaling config data: %v", err)
	}
	return config
}
