package main

import (
	"log"

	"ClientService/internal/application"
	Settings "ClientService/internal/config"
	"ClientService/internal/logging"
)

func main() {
	log.Println("Initialising client application...")
	log.Println("Collecting config...")
	cfg := Settings.GetConfig()
	log.Println("Getting logger of client application...")
	logger := logging.StartLog(cfg.LoggerConfig.LogLevel, cfg.LoggerConfig.LogFile)
	logger.Info().Msg("Starting client application...")
	app, err := application.NewApplication(cfg, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Application can't start")
	}
	app.Run()
}
