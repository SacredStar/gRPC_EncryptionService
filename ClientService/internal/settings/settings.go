package Settings

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog"
	"sync"
)

type ClientConfig struct {
	HomeDir  string `env:"HOME"`
	Port     int    `env:"PORT" envDefault:"3000"`
	Password string `env:"PASSWORD,unset"`

	Hosts          []string `env:"HOSTS" envSeparator:":"`
	TimeoutContext int      `env:"TIMEOUT" envDefault:"10"` // * time.Seconds
	TempFolder     string   `env:"TEMP_FOLDER" envDefault:"${HOME}/tmp" envExpand:"true"`
	LoggerConfig   struct {
		LogLevel     zerolog.Level `env:"LOGLEVEL" envDefault:"-1"`
		IsProduction bool          `env:"PRODUCTION"`
		LogFile      string        `env:"LOG" envDefault:"./ClientService/internal/logging/log.txt"`
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
