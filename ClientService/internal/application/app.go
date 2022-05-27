package application

import (
	_ "ClientService/docs"
	"ClientService/internal/logging"
	"ClientService/internal/metrics"
	Settings "ClientService/internal/settings"
	"context"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
	"net"
	"net/http"
	"time"
)

type Application struct {
	cfg        *Settings.ClientConfig
	logger     *logging.Logger
	router     *httprouter.Router
	httpServer *http.Server
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

	// TODO: CORS?

	logger.Info().Msg("heartbeat metric initializing...")
	metricHandler := metrics.Handler{}
	metricHandler.Register(router)
	return router
}

func (app *Application) Run() {
	app.startHTTP()
}

func (app *Application) startHTTP() {
	app.logger.Info().Msg("starting HTTP...")

	var listener net.Listener
	app.logger.Info().Msgf("binding application to host: %s and port: %s", app.cfg.Host, app.cfg.Port)
	var err error
	listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", app.cfg.Host, app.cfg.Port))
	if err != nil {
		app.logger.Fatal().Err(err).Msg("Can't start listening")
	}
	// TODO: CORS enable?
	/*c := cors.New(cors.Options{
		AllowedMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodOptions, http.MethodDelete},
		AllowedOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"Location", "Charset", "Access-Control-Allow-Origin", "Content-Type", "content-type", "Origin", "Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token"},
		OptionsPassthrough: true,
		ExposedHeaders:     []string{"Location", "Authorization", "Content-Disposition"},
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})

	handler := c.Handler(app.router)

	app.httpServer = &http.Server{
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}*/
	handler := app.router
	app.httpServer = &http.Server{
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	app.logger.Info().Msg("Application initialise completely and started.")

	if err := app.httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			app.logger.Warn().Msg("server shutdown.")
		default:
			app.logger.Fatal().Err(err).Msg("Can't start Server.")
		}
	}
	err = app.httpServer.Shutdown(context.Background())
	if err != nil {
		app.logger.Fatal().Err(err).Msg("Shutdown problem.")
	}
}
