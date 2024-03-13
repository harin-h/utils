package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	SecretKey string `yaml:"secret_key"`
	Password  string `yaml:"password"`
}

var config *Config

func init() {
	err := readConfig()
	if err != nil {
		panic(err)
	}
}

func readConfig() error {
	configObj := &Config{}
	configFile, err := os.ReadFile("./config.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(configFile, configObj)
	if err != nil {
		return err
	}
	config = configObj
	return nil
}
