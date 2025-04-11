package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Struct untuk memetakan data YAML
type ConfigYaml struct {
	App      ConfigApp      `yaml:"app"`
	Database ConfigDatabase `yaml:"database"`
}

type ConfigApp struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}

type ConfigDatabase struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
}

var (
	Cfg *ConfigYaml
)

func init() {
	file, err := os.Open("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(&Cfg); err != nil {
		log.Fatal(err)
	}

	log.Println("[CONFIG] Get config Success")
}

func GetAppConfig() ConfigApp {
	return Cfg.App
}
