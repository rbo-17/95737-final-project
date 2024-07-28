package cassandra

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
		Username: os.Getenv("CASSANDRA_USERNAME"),
		Password: os.Getenv("CASSANDRA_PASSWORD"),
		Address:  os.Getenv("CASSANDRA_ADDRESS"),
		DbName:   os.Getenv("MONGODB_DB_NAME"),
	}

	if c.Username == "" {
		return c, errors.New("CASSANDRA_USERNAME env var not set")
	}

	if c.Password == "" {
		return c, errors.New("CASSANDRA_PASSWORD env var not set")
	}

	if c.Address == "" {
		return c, errors.New("CASSANDRA_ADDRESS env var not set")
	}

	if c.DbName == "" {
		return c, errors.New("CASSANDRA_DB_NAME env var not set")
	}

	return c, nil
}
