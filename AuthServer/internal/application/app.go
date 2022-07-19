package application

import (
	"AuthServer/internal/config"
	"AuthServer/internal/logging"
	"crypto/sha256"
	"fmt"
	"net"
	"net/http"
)

type Application struct {
	cfg        *config.AuthServerConfig
	logger     *logging.Logger
	httpServer *http.Server
}

func NewAuthApplication(cfg *config.AuthServerConfig, logger *logging.Logger) (Application, error) {
	logger.Info().Msg("Starting application...")
	app := Application{
		cfg:    cfg,
		logger: logger,
	}
	app.logger.Info().Msg("Auth Application correctly created")
	return app, nil
}

func (app *Application) Run() {
	listener, _ := net.Listen("tcp", app.cfg.Host+":"+app.cfg.Port) // открываем слушающий сокет
	app.logger.Info().Msg("Starting listen TCP connection...")
	for {
		app.logger.Info().Msg("New TCP connection received...")
		conn, err := listener.Accept() // принимаем TCP-соединение от клиента и создаем новый сокет
		if err != nil {
			continue
		}
		app.logger.Info().Msg("Handling Connection...")
		go handleClient(conn) // обрабатываем запросы клиента в отдельной го-рутине
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close() // закрываем сокет при выходе из функции

	buf := make([]byte, 32) // буфер для чтения клиентских данных
	for {
		_, err := conn.Read(buf) // читаем из сокета
		if err != nil {
			fmt.Println(err)
			break
		}
		//TODO: testing with hasher function
		hasher := sha256.New()
		hasher.Write(buf)

		conn.Write(hasher.Sum(nil)) // пишем в сокет
	}
}
