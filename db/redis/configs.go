package redis

import (
	"errors"
	"os"
)

type Configs struct {
	Password string
	Address  string
}

func GetRedisConfigs() (Configs, error) {
	c := Configs{
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
