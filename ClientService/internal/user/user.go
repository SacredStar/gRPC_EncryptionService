package user

import (
	"crypto"
)

type User struct {
	UID      string
	username string
	password string
	Auth
	publicKey  crypto.PublicKey
	privateKey crypto.PrivateKey
}

func (u User) GetStorage() bool {
	//TODO: implement me
	panic("implement me")
}

func (u User) AddUpdateStorageRecord() bool {
	//TODO: implement me
	panic("implement me")
}

func (u User) DeleteRecord() bool {
	//TODO: implement me
	panic("implement me")
}

func (u *User) SetUserName(newname string) {
	u.username = newname
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u User) GetUserName() string {
	return u.username
}

func (u User) GetPassword() string {
	return u.password
}
