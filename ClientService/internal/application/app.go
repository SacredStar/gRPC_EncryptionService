package application

import (
	"ClientService/internal/user"
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	_ "ClientService/docs"
	Settings "ClientService/internal/config"
	guihtml "ClientService/internal/gui"
	"ClientService/internal/logging"
	"ClientService/internal/metrics"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Application struct {
	cfg        *Settings.ClientConfig
	logger     *logging.Logger
	router     *httprouter.Router
	httpServer *http.Server
	User       *user.User
}

func NewApplication(cfg *Settings.ClientConfig, logger *logging.Logger) (Application, error) {
	logger.Info().Msg("Starting application...")
	logger.Info().Msg("Starting routing...")
	//Установим значения из конфига в структуру User
	//TODO: collides with package name
	user := SetUserConfiguration(cfg)
	router := startRouting(logger, cfg, user)
	app := Application{
		cfg:    cfg,
		logger: logger,
		router: router,
	}
	app.logger.Info().Msg("Application correctly created")
	return app, nil
}

func startRouting(logger *logging.Logger, cfg *Settings.ClientConfig, user *user.User) *httprouter.Router {
	logger.Info().Msg("Router initialising...")
	router := httprouter.New()

	logger.Info().Msg("Swagger initialising...")
	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("swagger/index.html", http.StatusMovedPermanently))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	logger.Info().Msg("GUI initialising...")

	guiHandler := guihtml.Handler{
		HtmlRoot: cfg.HTMLRootFolder,
		User:     user,
		Logger:   logger,
	}

	guiHandler.Register(router)

	logger.Info().Msg("Metrics initializing...")
	metricHandler := metrics.Handler{}
	metricHandler.Register(router)
	return router
}

func SetUserConfiguration(cfg *Settings.ClientConfig) *user.User {
	User := user.User{}
	User.SetUserName(cfg.Login)
	User.SetPassword(cfg.Password)
	User.SetServerPort(cfg.AuthConfig.AuthPort)
	User.SetServerHost(cfg.AuthConfig.AuthHostName)
	return &User
}

func (app *Application) Run() {
	app.startHTTP(app.cfg.IsCORSEnabled)
}

func (app *Application) startHTTP(IsCORSEnabled bool) {
	app.logger.Info().Msg("Starting HTTP...")

	var listener net.Listener
	app.logger.Info().Msgf("Binding application to host: %s and port: %s", app.cfg.Host, app.cfg.Port)
	var err error
	listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", app.cfg.Host, app.cfg.Port))
	if err != nil {
		app.logger.Fatal().Err(err).Msg("Can't start listening")
	}
	var handler http.Handler
	// TODO: Understood CORS methods...
	if IsCORSEnabled {
		c := cors.New(cors.Options{
			AllowedMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodOptions, http.MethodDelete},
			AllowedOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
			AllowCredentials:   true,
			AllowedHeaders:     []string{"Location", "Charset", "Access-Control-Allow-Origin", "Content-Type", "content-type", "Origin", "Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token"},
			OptionsPassthrough: true,
			ExposedHeaders:     []string{"Location", "Authorization", "Content-Disposition"},
			// Enable Debugging for testing, consider disabling in production
			Debug: false,
		})
		handler = c.Handler(app.router)
	} else {
		handler = app.router
	}
	app.httpServer = &http.Server{
		Handler:      handler,
		WriteTimeout: time.Duration(app.cfg.TimeoutContext) * time.Second,
		ReadTimeout:  time.Duration(app.cfg.TimeoutContext) * time.Second,
	}

	app.logger.Info().Msg("Application initialise completely and started.")

	if err := app.httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			app.logger.Warn().Err(err).Msg("Server shutdown.")
		default:
			app.logger.Fatal().Err(err).Msg("Can't start Server.")
		}
	}
	err = app.httpServer.Shutdown(context.Background())
	if err != nil {
		app.logger.Fatal().Err(err).Msg("Shutdown problem.")
	}
}
