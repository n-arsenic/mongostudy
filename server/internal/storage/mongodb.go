package storage

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBConfig struct {
	Host     string `env:"DB_HOST,required"`
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	DBName   string `env:"DB_NAME,required"`
}

func (c DBConfig) Validate() error {
	return nil
}

func (c DBConfig) ConnectionURL() string {
	return "mongodb://" + c.User + ":" + c.Password + "@" + c.Host
}

const (
	RequestsTimeLimit   = 3 * time.Second
	ConnectionTimeLimit = 10 * time.Second
)

type MongoDB struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoDB(ctx context.Context, conf DBConfig) (*MongoDB, error) {
	if err := conf.Validate(); err != nil {
		return nil, fmt.Errorf("db parameters is incorrect: %v", err)
	}
	opt := options.Client()
	opt.ApplyURI(conf.ConnectionURL())
	opt.SetConnectTimeout(ConnectionTimeLimit)
	client, err := mongo.NewClient(opt)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &MongoDB{
		client: client,
		db:     client.Database(conf.DBName),
	}, nil
}

func (m *MongoDB) Close() {
	if err := m.client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func (m *MongoDB) GetCollection(name string) *mongo.Collection {
	return m.db.Collection(name)
}
