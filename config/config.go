package config

import (
	"github.com/georgizhivankin/go-slacking-cow/helpers"
)

// Define a struct that holds all config variables for the cow app
type ConfigVariables struct {
	ApiName     string
	ApiUrl      string
	ApiMethod   string
	ApiFormat   string
	ApiKey      int
	ApiLanguage string
}

func NewConfigVariables() *ConfigVariables {
	return &ConfigVariables{
		ApiName:     "Forismatic",
		ApiUrl:      "http://api.forismatic.com/api/1.0/",
		ApiMethod:   "getQuote",
		ApiFormat:   "json",
		ApiKey:      helpers.GenerateRandomNumber(6),
		ApiLanguage: "en"}
}
