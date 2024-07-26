// Helpful docs:
// - Taken from: https://stackoverflow.com/a/21112176
// - https://go.dev/doc/database/prepared-statements

package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/rbo-17/95737-final-project/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var ctx context.Context

func init() {
	ctx = context.Background()
}

type KeyValue struct {
	Key   string
	Value []byte
}

type MySQL struct {
	Name      string
	TableName string
	C         *Configs
	Db        *sql.DB
}

func NewMySQL() *MySQL {
	return &MySQL{
		Name:      utils.DbNameMySQL,
		TableName: "KEY_VALUE",
		C:         nil,
		Db:        nil,
	}
}

func (m *MySQL) Init() error {
	c, err := GetConfigs()
	if err != nil {
		return err
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", c.Username, c.Password, c.Address, c.DbName)

	// See setup documentation here: https://github.com/go-sql-driver/mysql
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)

	m.C = &c
	m.Db = db

	return nil
}

func (m *MySQL) GetName() string {
	return m.Name
}

func (m *MySQL) GetKey(keyId string) string {
	return keyId
}

func (m *MySQL) Get(k string) ([]byte, error) {

	stmt, err := m.Db.Prepare(fmt.Sprintf("SELECT * FROM %s WHERE ID = ?", m.TableName))
	if err != nil {
		return nil, err
	}

	var res KeyValue
	err = stmt.QueryRow(k).Scan(&res.Key, &res.Value)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return res.Value, nil
}

func (m *MySQL) Put(k string, v []byte) error {

	input := make(map[string][]byte, 1)
	input[k] = v

	err := m.PutMany(input)
	if err != nil {
		return err
	}

	return nil
}

func (m *MySQL) PutMany(kv map[string][]byte) error {

	sqlStr := fmt.Sprintf("INSERT INTO %s(ID, Value) VALUES", m.TableName)
	var vals []interface{}
	for k, v := range kv {
		sqlStr += " (?, ?),"
		vals = append(vals, k, v)
	}

	// trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]

	// add semicolon
	sqlStr += ";"

	// prepare the statement
	stmt, err := m.Db.Prepare(sqlStr)
	if err != nil {
		return err
	}

	// format all vals at once
	_, err = stmt.Exec(vals...)
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}

func (m *MySQL) DeleteAll() error {
	_, err := m.Db.Query(fmt.Sprintf("DELETE FROM %s;", m.TableName))
	if err != nil {
		return err
	}

	return nil
}
