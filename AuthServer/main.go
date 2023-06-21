package main

import (
	"AuthServer/internal/application"
	"AuthServer/internal/config"
	"AuthServer/internal/logging"
	"log"
)

func main() {
	log.Println("Starting Auth Server...")
	log.Println("Collecting config...")
	cfg := config.GetConfig()
	log.Println("Getting logger of Auth application...")
	logger := logging.StartLog(cfg.LoggerConfig.LogLevel, cfg.LoggerConfig.LogFile)
	logger.Info().Msg("Starting Auth application...")
	app, err := application.NewAuthApplication(cfg, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Application can't start")
	}
	app.Run()

}
