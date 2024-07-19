package redis

import (
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

type RedisConfigs struct {
	Password string
	Address  string
}

func GetRedisConfigs() (RedisConfigs, error) {
	c := RedisConfigs{
		Password: os.Getenv("REDIS_PASSWORD"),
		Address:  os.Getenv("REDIS_ADDRESS"),
	}

	if c.Password == "" {
		return c, errors.New("REDIS_PASSWORD env var not set")
	}

	if c.Address == "" {
		return c, errors.New("REDIS_ADDRESS env var not set")
	}

	return c, nil
}

func GetRedisInstance() (*redis.Client, error) {

	c, err := GetRedisConfigs()
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", c.Address),
		Password: c.Password,
		DB:       0, // use default DB
	})

	return rdb, nil
}
