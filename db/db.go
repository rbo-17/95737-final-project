package db

type Db interface {
	Init() error
	GetKey(kid string) string
	GetName() string
	Get(k string) (*[]byte, error)
	Put(k string, v *[]byte) error
	PutMany(kv map[string]*[]byte) error
	DeleteAll() error
	Close() error
}
