package application

import (
	"AuthServer/internal/config"
	"AuthServer/internal/logging"
	"bufio"
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

//TODO: revork to secure connection? grpc? tls?
func handleClient(conn net.Conn) {
	//TODO: переделать на нормальную обработку логина пароля от клиента
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		//TODO: testing with hasher function,revork to interface with some CSP's?
		fmt.Printf("Печатает буфер:%X\n", message)
		hasher := sha256.New()
		hasher.Write([]byte(message))
		result := hasher.Sum(nil)
		fmt.Printf("Печатаем полученный хеш:%X\n", result)
		//TODO: Тут должна быть пересылка на сервер, возможно стоит обработать и логин/пароль?
		if _, err := conn.Write(result); err != nil {
			fmt.Println("Error,cant send token to connection")
		} // пишем в сокет
		if err := conn.Close(); err != nil {
			fmt.Println("Error,cant close connection")
		}
		break
	}
}
