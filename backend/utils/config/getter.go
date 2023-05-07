package config

import (
	"errors"
	"io/ioutil"
)

func DatabaseConfig() (*DatabaseConf, error) {
	if config == nil {
		return nil, errors.New("config is not initialised")
	}
	return &config.Database, nil
}

func ServerConfig() (*ServerConf, error) {
	if config == nil {
		return nil, errors.New("config is not initialised")
	}
	return &config.Server, nil
}

func S3Config() (*S3Conf, error) {
	if config == nil {
		return nil, errors.New("config is not initialised")
	}
	return &config.S3, nil
}

func GetSecret(path string) (string, error) {
	secret, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(secret), nil
}
