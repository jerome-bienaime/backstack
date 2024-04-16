package config

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type EnvConfig struct {
	SQLITE_PATH      string `validate:"required"`
	SQLITE_DEBUG     string `validate:"required"`
	HTTP_PORT        string `validate:"required"`
	HTTP_CACHE       string `validate:"required"`
	HTTP_COMPRESSION string `validate:"required"`
	HTTP_SSL         string `validate:"required"`
	HTTP_TIMEOUT     string `validate:"required"`
	MODE             string `validate:"required"`
	SECRET_KEY       string `validate:"required"`
}

var ENV = [4]string{
	"TEST",
	"DEV",
	"STAGING",
	"PROD",
}

func ReadEnvFile(envName string) (*EnvConfig, error) {
	var currentEnv string
	for _, name := range ENV {
		if name == strings.ToUpper(envName) {
			currentEnv = strings.ToLower(name)
		}
	}
	if currentEnv == "" {
		return &EnvConfig{}, errors.New("could not file env name" + envName)
	}
	viper.SetConfigFile("." + currentEnv + ".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return &EnvConfig{}, err
	}

	var c EnvConfig
	err := viper.Unmarshal(&c)
	if err != nil {
		return &EnvConfig{}, err
	}

	validator := validator.New()
	if err := validator.Struct(c); err != nil {
		return &EnvConfig{}, err
	}

	return &c, nil
}
