package db

type Db interface {
	Init()
	GetKey(kid string) string
	GetName() string
	Get(k string) ([]byte, error)
	Put(k string, v []byte) error
	//PutMany()
	DeleteAll() error
}
