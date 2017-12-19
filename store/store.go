package store

type Store interface {
	Initialize()
	Read(key string) error
	Write(key, val string) error
	Start() error
	Commit() error
	Abort() error
}
