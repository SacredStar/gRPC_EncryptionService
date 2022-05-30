package user

import "crypto"

type user struct {
	UID        string
	username   string
	password   string
	auth       auth
	publicKey  crypto.PublicKey
	privateKey crypto.PrivateKey
}

func (u user) GetStorage() bool {
	//TODO implement me
	panic("implement me")
}

func (u user) AddUpdateStorageRecord() bool {
	//TODO implement me
	panic("implement me")
}

func (u user) DeleteRecord() bool {
	//TODO implement me
	panic("implement me")
}
