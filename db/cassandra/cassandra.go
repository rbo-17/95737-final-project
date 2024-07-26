package cassandra

import (
	"context"
	"github.com/rbo-17/95737-final-project/utils"
	"github.com/redis/go-redis/v9"
)

var ctx context.Context

func init() {
	ctx = context.Background()
}

type Cassandra struct {
	Name string
	C    *Configs
	Rdb  *redis.Client
}

func NewCassandra() *Cassandra {
	return &Cassandra{
		Name: utils.DbNameCassandra,
		C:    nil,
		Rdb:  nil,
	}
}

func (c *Cassandra) Init() error {

}

func (c *Cassandra) GetName() string {
	return r.Name
}

func (c *Cassandra) GetKey(keyId string) string {
	return keyId
}

func (c *Cassandra) Get(k string) ([]byte, error) {

}

func (c *Cassandra) Put(k string, v []byte) error {

}

func (c *Cassandra) PutMany(kv map[string][]byte) error {

}

func (c *Cassandra) DeleteAll() error {

}
