package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"
)

type MysqlConfigs struct {
	Username string
	Password string
	Address  string
	DbName   string
}

func GetMysqlConfigs() (MysqlConfigs, error) {
	c := MysqlConfigs{
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

func GetMysqlInstance() (*sql.DB, error) {
	c, err := GetMysqlConfigs()
	if err != nil {
		panic(err)
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", c.Username, c.Password, c.Address, c.DbName)

	// See setup documentation here: https://github.com/go-sql-driver/mysql
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)

	return db, nil
}
