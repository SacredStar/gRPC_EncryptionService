package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"sync"

	"github.com/rs/zerolog"
)

type AuthServerConfig struct {
	HomeDir        string `env:"HOME"`
	Port           string `env:"PORT" envDefault:"2000"`
	Host           string `env:"HOST" envDefault:"0.0.0.0"`
	TimeoutContext int    `env:"TIMEOUT" envDefault:"10"` // * time.Seconds
	TempFolder     string `env:"TEMP_FOLDER" envDefault:"${HOME}/tmp" envExpand:"true"`

	LoggerConfig struct {
		LogLevel     zerolog.Level `env:"LOGLEVEL" envDefault:"-1"`
		IsProduction bool          `env:"PRODUCTION"`
		LogFile      string        `env:"LOG" envDefault:"./internal/logging/log.txt"`
	}
}

var instance *AuthServerConfig
var once sync.Once

func GetConfig() *AuthServerConfig {
	once.Do(func() {
		instance = &AuthServerConfig{}
		if err := env.Parse(instance); err != nil {
			fmt.Printf("%+v\n", err)
		}

	})
	return instance
}
