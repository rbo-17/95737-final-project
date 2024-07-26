package mysql

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
		Username: os.Getenv("MYSQL_USERNAME"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Address:  os.Getenv("MYSQL_ADDRESS"),
		DbName:   os.Getenv("MYSQL_DB_NAME"),
	}

	if c.Username == "" {
		return c, errors.New("MYSQL_USERNAME env var not set")
	}

	if c.Password == "" {
		return c, errors.New("MYSQL_PASSWORD env var not set")
	}

	if c.Address == "" {
		return c, errors.New("MYSQL_ADDRESS env var not set")
	}

	if c.DbName == "" {
		return c, errors.New("MYSQL_DB_NAME env var not set")
	}

	return c, nil
}
