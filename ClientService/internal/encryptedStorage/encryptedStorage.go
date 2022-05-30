package encryptedStorage

// storager TODO: Change current stub
type storager interface {
	GetStorage() bool
	AddUpdateStorageRecord() bool
	DeleteRecord() bool
}

type record struct {
	site     string
	login    string
	password string // TODO: Encrypted?
}

type storage struct {
	record []record
}
