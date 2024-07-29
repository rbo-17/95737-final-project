// See connection guide at: https://www.mongodb.com/docs/drivers/go/current/fundamentals/connections/connection-guide/#std-label-golang-connection-guide
package mongodb

import (
	"context"
	"fmt"
	"github.com/rbo-17/95737-final-project/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx context.Context

func init() {
	ctx = context.Background()
}

type KeyValue struct {
	Key   string
	Value []byte
}

type MongoDB struct {
	Name       string
	C          *Configs
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewMongoDB() *MongoDB {
	return &MongoDB{
		Name:       utils.DbNameMongoDB,
		Client:     nil,
		Collection: nil,
	}
}

func (m *MongoDB) Init() error {

	c, err := GetConfigs()
	if err != nil {
		return err
	}
	m.C = &c

	protocol := "mongodb"
	userName := c.Username
	password := c.Password
	hostName := c.Address
	connOpts := c.DbName

	uri := fmt.Sprintf("%s://%s:%s@%s:27017/%s", protocol, userName, password, hostName, connOpts)

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return err
	}

	m.Client = client
	m.Collection = m.Client.Database("95737").Collection("final-project")

	return nil
}

func (m *MongoDB) GetKey(kid string) string {
	return kid
}

func (m *MongoDB) GetName() string {
	return m.Name
}

func (m *MongoDB) Get(k string) (*[]byte, error) {

	filter := bson.D{{"_id", k}}

	res := KeyValue{}
	err := m.Collection.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res.Value, nil
}

func (m *MongoDB) Put(k string, v *[]byte) error {

	record := bson.D{
		{"_id", k},
		{"value", *v},
	}

	_, err := m.Collection.InsertOne(ctx, record)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoDB) PutMany(kv map[string]*[]byte) error {
	var input []interface{}
	for k, v := range kv {
		record := bson.D{
			{"_id", k},
			{"value", *v},
		}

		input = append(input, record)
	}

	_, err := m.Collection.InsertMany(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoDB) DeleteAll() error {
	err := m.Collection.Drop(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoDB) Close() error {

	err := m.Client.Disconnect(ctx)
	if err != nil {
		return err
	}

	return nil
}
