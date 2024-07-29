package redis

import (
	"context"
	"fmt"
	"github.com/rbo-17/95737-final-project/utils"
	"github.com/redis/go-redis/v9"
)

var ctx context.Context

func init() {
	ctx = context.Background()
}

type Redis struct {
	Name string
	C    *Configs
	Rdb  *redis.Client
}

func NewRedis() *Redis {
	return &Redis{
		Name: utils.DbNameRedis,
		C:    nil,
		Rdb:  nil,
	}
}

func (r *Redis) Init() error {
	c, err := GetRedisConfigs()
	if err != nil {
		return err
	}
	r.C = &c

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", c.Address),
		Password: c.Password,
		DB:       0, // use default DB
	})

	r.Rdb = rdb

	return nil
}

func (r *Redis) GetName() string {
	return r.Name
}

func (r *Redis) GetKey(keyId string) string {
	return fmt.Sprintf(fmt.Sprintf("smtxt:%s", keyId))
}

func (r *Redis) Get(k string) (*[]byte, error) {
	value, err := r.Rdb.Get(ctx, k).Bytes()
	if err != nil {
		return &[]byte{}, err
	}

	return &value, nil
}

func (r *Redis) Put(k string, v *[]byte) error {
	err := r.Rdb.Set(ctx, k, *v, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Redis) PutMany(kv map[string]*[]byte) error {

	var input []interface{}
	for k, v := range kv {
		input = append(input, k, *v)
	}

	err := r.Rdb.MSet(ctx, input...).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Redis) DeleteAll() error {
	err := r.Rdb.FlushDB(ctx).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Redis) Close() error {

	err := r.Rdb.Close()
	if err != nil {
		return err
	}

	return nil
}
