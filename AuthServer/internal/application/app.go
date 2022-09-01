package application

import (
	"AuthServer/internal/User"
	"AuthServer/internal/config"
	"AuthServer/internal/logging"
	"bufio"
	"fmt"
	GOST "github.com/SacredStar/Encryption.go/GostCrypto"
	"golang.org/x/sys/windows"
	"net"
	"net/http"
	"strings"
)

type Application struct {
	cfg        *config.AuthServerConfig
	logger     *logging.Logger
	httpServer *http.Server
	hProv      windows.Handle
}

func NewAuthApplication(cfg *config.AuthServerConfig, logger *logging.Logger) (Application, error) {
	logger.Info().Msg("Starting application...")
	logger.Info().Msg("Acquiring crypto context...")
	var hProvHandle windows.Handle
	if err := windows.CryptAcquireContext(&hProvHandle, nil, nil, 80, windows.CRYPT_VERIFYCONTEXT); err != nil {
		logger.Fatal().Msg("Cant AcquireContext of provider...")
		return Application{}, err
	}
	app := Application{
		cfg:    cfg,
		logger: logger,
		hProv:  hProvHandle,
	}
	app.logger.Info().Msg("Auth Application correctly created.")
	return app, nil
}

func (app *Application) Run() {
	listener, _ := net.Listen("tcp", app.cfg.Host+":"+app.cfg.Port) // открываем слушающий сокет
	app.logger.Info().Msg("Starting listen TCP connection...")
	for {

		conn, err := listener.Accept() // принимаем TCP-соединение от клиента и создаем новый сокет
		app.logger.Info().Msg("New TCP connection received...")
		if err != nil {
			continue
		}
		app.logger.Info().Msg("Handling Connection: " + conn.RemoteAddr().String() + "...")
		go handleClient(conn, app) // обрабатываем запросы клиента в отдельной го-рутине
	}
}

//TODO: revork to secure connection? grpc? tls?
//TODO: передлать на получение вначале определенного флага для действия?
func handleClient(conn net.Conn, app *Application) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		NewUser := processAuthMessage(message, err, app)
		app.logger.Info().Msgf("Печатаем полученный хеш:%X\n", NewUser)
		//TODO: Тут должна быть пересылка на сервер?
		//TODO: перед пересылкой необходимо удостовериться что такой User/password имеется в базе данных
		if _, err := conn.Write([]byte(NewUser.Token)); err != nil {
			app.logger.Fatal().Msg("Error,cant send token to connection")
		} // пишем в сокет
		if err := conn.Close(); err != nil {
			app.logger.Info().Msg("Error,cant close connection")
		}
		break
	}
}

// processAuthMessage - получаем логин пароль из сообщения,считаем хеш и закидываем его в структуру
func processAuthMessage(message string, err error, app *Application) *User.User {
	//need to trim message from last new line char
	message = strings.TrimSuffix(message, "\n")
	splitedMessage := strings.Split(message, " ")
	str := []byte(message)
	result, err := GOST.CreateHashFromData(app.hProv, GOST.CALG_GR3411_2012_256, &str[0], uint32(len(str)))
	if err != nil {
		app.logger.Fatal().Msgf("error. Can't create hash from message:%s", str)
	}
	userFromConnection := &User.User{
		Login:    splitedMessage[0],
		Password: splitedMessage[1],
		Token:    string(result),
	}
	return userFromConnection
}
