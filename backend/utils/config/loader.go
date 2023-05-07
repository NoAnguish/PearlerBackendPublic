package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var config *Config

func LoadConfig(path string) error {
	config = &Config{}

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	dec := yaml.NewDecoder(f)
	err = dec.Decode(config)

	if err != nil {
		config = nil
		return err
	}
	return nil
}
