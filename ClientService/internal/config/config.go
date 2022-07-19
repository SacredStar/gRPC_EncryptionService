package Config

import (
	"fmt"
	"sync"

	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog"
)

type ClientConfig struct {
	HomeDir        string `env:"HOME"`
	Port           string `env:"PORT" envDefault:"3000"`
	Login          string `env:"LOGIN" envDefault:"LOGIN"`
	Password       string `env:"PASSWORD,unset"`
	HTMLRootFolder string `env:"HTML_ROOT" envDefault:"./internal/gui/main.html"`

	Host           string `env:"HOST" envDefault:"0.0.0.0"`
	TimeoutContext int    `env:"TIMEOUT" envDefault:"10"` // * time.Seconds
	IsCORSEnabled  bool   `env:"CORS" envDefault:"false"`
	TempFolder     string `env:"TEMP_FOLDER" envDefault:"${HOME}/tmp" envExpand:"true"`

	LoggerConfig struct {
		LogLevel     zerolog.Level `env:"LOGLEVEL" envDefault:"-1"`
		IsProduction bool          `env:"PRODUCTION"`
		LogFile      string        `env:"LOG" envDefault:"./internal/logging/log.txt"`
	}
	AuthConfig struct {
		AuthHostName string `env:"AUTHHOST" envDefault:"0.0.0.0"`
		AuthPort     string `env:"AUTHPORT" envDefault:"3000"`
	}
}

var instance *ClientConfig
var once sync.Once

func GetConfig() *ClientConfig {
	once.Do(func() {
		instance = &ClientConfig{}
		if err := env.Parse(instance); err != nil {
			fmt.Printf("%+v\n", err)
		}

	})
	return instance
}
