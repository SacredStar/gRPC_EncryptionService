package main

import (
	"log"

	"ClientService/internal/application"
	"ClientService/internal/logging"
	Settings "ClientService/internal/settings"
)

/*var (
	addr = fmt.Sprintf("%s", "localhost:9999")
)*/

func main() {
	log.Println("Starting client application...")
	log.Println("Collecting config...")
	cfg := Settings.GetConfig()
	log.Println("Getting logger of client application...")
	logger := logging.StartLog(cfg.LoggerConfig.LogLevel, cfg.LoggerConfig.LogFile)
	logger.Info().Msg("Starting application...")
	app, err := application.NewApplication(cfg, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Application can't start")
	}
	app.Run()
	//logger.Info().Str("foo", "bar").Msg("Hello world")
	//StartgRPC()
}
