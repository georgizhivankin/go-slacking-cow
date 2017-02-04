package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Define a struct that holds all global configuration variables for the cow app
type Specification struct {
	AppName    string `envconfig:"APP_NAME" default:"Cowsay implementation for Slack"`
	AppVersion string `envconfig:"APP_VERSION" default:"1.0"`
	Debug      bool   `envconfig:"DEBUG"`
	LogLevel   string `envconfig:"LOG_LEVEL" default:"info"`
	Database   Database
	QuotesApi  QuotesApi
}

type Database struct {
	DSN  string `envconfig:"DATABASE_DSN" default:"mongo"`
	Port string `envconfig:"DATABASE_PORT" default:"27017"`
	Name string `envconfig:"DATABASE_NAME" default:"local"`
}

type QuotesApi struct {
	ApiName     string `envconfig:"QUOTES_API_NAME" default:"Forismatic"`
	ApiUrl      string `envconfig:"QUOTES_API_URL" default:"http://api.forismatic.com/api/1.0/"`
	ApiMethod   string `envconfig:"QUOTES_API_METHOD" default:"getQuote"`
	ApiFormat   string `envconfig:"QUOTES_API_FORMAT" default:"json"`
	ApiKey      int    `envconfig:"QUOTES_API_KEY" default:"3594941"`
	ApiLanguage string `envconfig:"QUOTES_API_LANGUAGE" default:"en"`
}

//LoadEnv loads environment variables
func LoadEnv() (*Specification, error) {
	var config Specification
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
