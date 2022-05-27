package application

import (
	"ClientService/internal/logging"
	Settings "ClientService/internal/settings"
)

type Application struct {
	cfg    *Settings.ClientConfig
	logger *logging.Logger
}

func NewApplication(cfg *Settings.ClientConfig, logger *logging.Logger) (Application, error) {
	logger.Info().Msg("Started application with cfg and logger")
	app := Application{
		cfg:    cfg,
		logger: logger,
	}

	return app, nil
}

/*

}
*/
