package application

import (
	"gRPCClientServerForEncryption/ClientService/internal/logging"
	Settings "gRPCClientServerForEncryption/ClientService/internal/settings"
)

type Application struct {
}

func NewApplication(cfg *Settings.ClientConfig, logger *logging.Logger) (*Application, error) {
	logger.Info().Msgf("Started application with cfg: %#v", cfg)
	return &Application{}, nil
}
