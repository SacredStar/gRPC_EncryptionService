package user

type authenticate interface {
	Auth() string
	GetToken() string
}

type auth struct {
	token    string
	algID    string
	provider string
}

func (a auth) Auth() string {
	//TODO: Send login password to auth server, return token
	return "token"
}

func (a auth) GetToken() string {
	return a.token
}
