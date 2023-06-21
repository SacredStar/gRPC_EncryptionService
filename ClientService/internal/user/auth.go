package user

import (
	"bufio"
	"fmt"
	"net"
)

type Auth struct {
	token      string
	ServerHost string
	ServerPort string
}

func (a *Auth) Authentificate(login, password string) {
	//TODO: Send login password to auth server, return token
	//TODO: set info from config?
	conn, err := net.Dial("tcp", "0.0.0.0:2000") // открываем TCP-соединение к серверу
	if err != nil {
		fmt.Println("error connecting to Auth server")
		return
	}
	for {
		data := []byte(login + " " + password + "\n")
		if _, err := conn.Write(data); err != nil {
			fmt.Println("Error sending login/password to auth server")
			return
		}
		token, _ := bufio.NewReader(conn).ReadBytes('.')
		a.token = string(token)
		if err := conn.Close(); err != nil {
			fmt.Println("Error,cant close connection")
		}
		break
	}
}

func (a *Auth) GetToken() string {
	return a.token
}

func (a *Auth) SetServerHost(host string) {
	a.ServerHost = host
}

func (a *Auth) SetServerPort(port string) {
	a.ServerPort = port
}
