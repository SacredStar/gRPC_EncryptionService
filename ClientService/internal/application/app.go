package application

import (
	"ClientService/internal/logging"
	"ClientService/internal/metrics"
	Settings "ClientService/internal/settings"
	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type Application struct {
	cfg    *Settings.ClientConfig
	logger *logging.Logger
	router *httprouter.Router
}

func NewApplication(cfg *Settings.ClientConfig, logger *logging.Logger) (Application, error) {
	logger.Info().Msg("Starting application...")
	logger.Info().Msg("Starting routing...")
	router := startRouting(logger)
	app := Application{
		cfg:    cfg,
		logger: logger,
		router: router,
	}
	app.logger.Info().Msg("Application correctly started")
	return app, nil
}

func startRouting(logger *logging.Logger) *httprouter.Router {
	logger.Info().Msg("Router initialising...")
	router := httprouter.New()

	logger.Info().Msg("Swagger initialising...")
	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("swagger/index.html", http.StatusMovedPermanently))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	logger.Info().Msg("heartbeat metric initializing...")
	metricHandler := metrics.Handler{}
	metricHandler.Register(router)
	return router
}
