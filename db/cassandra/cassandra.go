package cassandra

import (
	"context"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/rbo-17/95737-final-project/utils"
	"strings"
)

var ctx context.Context

func init() {
	ctx = context.Background()
}

type KeyValue struct {
	Key   string
	Value []byte
}

type Cassandra struct {
	Name      string
	TableName string
	C         *Configs
	Session   *gocql.Session
}

func NewCassandra() *Cassandra {
	return &Cassandra{
		Name:      utils.DbNameCassandra,
		TableName: strings.ToLower("KEY_VALUE"),
		C:         nil,
		Session:   nil,
	}
}

func (c *Cassandra) Init() error {
	configs, err := GetConfigs()
	if err != nil {
		return err
	}
	c.C = &configs

	cluster := gocql.NewCluster(fmt.Sprintf("%s:9042", c.C.Address))
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: c.C.Username,
		Password: c.C.Password,
	}
	cluster.Keyspace = c.C.DbName
	session, err := cluster.CreateSession()
	if err != nil {
		return err
	}

	c.Session = session

	return nil
}

func (c *Cassandra) GetName() string {
	return c.Name
}

func (c *Cassandra) GetKey(keyId string) string {
	return keyId
}

func (c *Cassandra) Get(k string) (*[]byte, error) {

	var v []byte
	res := c.Session.Query(fmt.Sprintf(`SELECT Payload FROM %s WHERE ID = ?;`, c.TableName), k).Iter()
	res.Scan(&v)

	if err := res.Close(); err != nil {
		return nil, err
	}

	return &v, nil
}

func (c *Cassandra) Put(k string, v *[]byte) error {
	err := c.Session.Query(fmt.Sprintf(`INSERT INTO %s (ID, Payload) VALUES (?, ?);`, c.TableName), k, *v).Exec()
	if err != nil {
		return err
	}

	return nil
}

func (c *Cassandra) PutMany(kv map[string]*[]byte) error {

	firstV := make([]byte, 0)
	for _, v := range kv {
		firstV = *v
	}

	if len(firstV) < 1000 {
		statement := "BEGIN BATCH "
		var args []interface{}
		for k, v := range kv {
			statement += fmt.Sprintf(`INSERT INTO %s (ID, Payload) VALUES (?, ?); `, c.TableName)
			args = append(args, k, v)
		}
		statement += " APPLY BATCH"

		res := c.Session.Query(statement, args...).Exec()
		if res != nil {
			return res
		}

	} else {
		for k, v := range kv {
			err := c.Put(k, v)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Cassandra) DeleteAll() error {
	err := c.Session.Query(fmt.Sprintf(`TRUNCATE %s;`, c.TableName)).Exec()
	if err != nil {
		return err
	}

	return nil
}

func (c *Cassandra) Close() error {
	c.Session.Close()

	return nil
}
