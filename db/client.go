package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type Client struct {
	ctx  context.Context
	c    *mongo.Client
	db   *mongo.Database
	coll map[string]*Collection
}

func NewClient(protocol string, host string, port int, dbname string) (*Client, error) {
	uri := fmt.Sprintf("%v://%v:%v", protocol, host, port)

	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri), options.Client().SetWriteConcern(writeconcern.Majority()))
	if err != nil {
		return nil, err
	}

	client := &Client{
		ctx:  context.Background(),
		c:    mongoClient,
		coll: make(map[string]*Collection),
	}

	err = client.Ping()
	if err != nil {
		return nil, err
	}

	mongoDB := mongoClient.Database(dbname)
	client.db = mongoDB

	return client, nil
}

func (c *Client) Ping() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	err = c.c.Ping(ctx, nil)
	return err
}

func (c *Client) Close() error {
	return c.c.Disconnect(c.ctx)
}

func (c *Client) Collection(collection string) *Collection {
	coll, ok := c.coll[collection]
	if ok && coll != nil {
		return coll
	}

	mongoColl := c.db.Collection(collection)

	coll = &Collection{mongoColl}
	c.coll[collection] = coll

	return coll
}
