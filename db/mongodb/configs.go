package mongodb

import (
	"errors"
	"os"
)

type Configs struct {
	Username string
	Password string
	Address  string
	DbName   string
}

func GetConfigs() (Configs, error) {
	c := Configs{
		Username: os.Getenv("MONGODB_USERNAME"),
		Password: os.Getenv("MONGODB_PASSWORD"),
		Address:  os.Getenv("MONGODB_ADDRESS"),
		DbName:   os.Getenv("MONGODB_DB_NAME"),
	}

	if c.Username == "" {
		return c, errors.New("MONGODB_USERNAME env var not set")
	}

	if c.Password == "" {
		return c, errors.New("MONGODB_PASSWORD env var not set")
	}

	if c.Address == "" {
		return c, errors.New("MONGODB_ADDRESS env var not set")
	}

	if c.DbName == "" {
		return c, errors.New("MONGODB_DB_NAME env var not set")
	}

	return c, nil
}
