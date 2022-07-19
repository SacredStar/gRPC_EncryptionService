package user

type authenticate interface {
	Auth()
	GetToken()
}

type Auth struct {
	token      string
	AlgID      string
	Provider   string
	ServerHost string
	ServerPort string
}

func (a *Auth) Authentificate(login, password string) {
	//TODO: Send login password to auth server, return token
	a.token = "TOKEN"
}

func (a *Auth) GetToken() string {
	return a.token
}
