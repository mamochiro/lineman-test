package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

// Configuration ...
type Configuration struct {
	// this section is set env param
	Port int    `env:"PORT" envDefault:"3000"`
	Env  string `env:"ENV" envDefault:"local"`

	CovidCaseURL string `env:"URL" envDefault:"http://static.wongnai.com/devinterview/covid-cases.json"`
}

// NewConfiguration ...
func NewConfiguration() Configuration {
	cfg := Configuration{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Errorf("%+v\n", err)
	}
	return cfg
}
