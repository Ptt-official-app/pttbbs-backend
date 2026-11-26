package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestNewClient(t *testing.T) {
	type args struct {
		protocol string
		host     string
		port     int
		dbname   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{protocol: "mongodb", host: "localhost", port: 27017, dbname: "testdb"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.protocol, tt.args.host, tt.args.port, tt.args.dbname)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			defer got.Close()
		})
	}
}

func TestClient_Ping_error(t *testing.T) {
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	require.NoError(t, err)

	client := &Client{c: mongoClient}
	require.NoError(t, client.Ping())

	// a disconnected client fails immediately, i.e. not through the context deadline.
	require.NoError(t, mongoClient.Disconnect(context.Background()))
	assert.Error(t, client.Ping())
}
