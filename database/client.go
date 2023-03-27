package database

import (
	"context"
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type client struct {
	mongoClient *mongo.Client
}

func New(host, username, password string, port uint64) (*client, error) {
	opts := options.Client()
	opts.SetAppName("go-mongo")
	opts.SetHosts([]string{host + ":" + strconv.FormatUint(port, 10)})
	opts.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	c, err := mongo.NewClient(opts)
	if err != nil {
		return nil, fmt.Errorf("error creating new database client: %v", err)
	}

	return &client{mongoClient: c}, nil
}

func (c *client) Connect(ctx context.Context) error {
	if err := c.mongoClient.Connect(ctx); err != nil {
		return err
	}

	if err := c.mongoClient.Ping(ctx, nil); err != nil {
		return err
	}
	defer c.mongoClient.Disconnect(ctx)

	return nil
}
