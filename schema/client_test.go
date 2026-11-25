package schema

import (
	"errors"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/sirupsen/logrus"
)

func TestUpdateClient(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer Client_c.Drop()

	client := NewClient("test_client_id2", types.CLIENT_TYPE_APP, nil, "localhost")

	type args struct {
		c *Client
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		expected *Client
	}{
		// TODO: Add test cases.
		{
			args:     args{client},
			expected: client,
		},
		{
			args:     args{client},
			wantErr:  true,
			expected: client,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UpdateClient(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateClient() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := GetClient("test_client_id2")
			logrus.Infof("UpdateClient: after got: e: %v", err)
			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
	}
}

func TestCreateClient(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer Client_c.Drop()

	client := NewClient("test_client_id3", types.CLIENT_TYPE_WEB, []string{"https://good.example/cb"}, "localhost")
	hijack := NewClient("test_client_id3", types.CLIENT_TYPE_WEB, []string{"https://evil.example/cb"}, "localhost")
	hijack.UpdateNanoTS = client.UpdateNanoTS + 1

	tests := []struct {
		name     string
		c        *Client
		wantErr  error
		expected *Client
	}{
		{name: "create", c: client, expected: client},
		{name: "existing client id is rejected and untouched", c: hijack, wantErr: ErrClientAlreadyExists, expected: client},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateClient(tt.c)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("CreateClient() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := GetClient("test_client_id3")
			if err != nil {
				t.Fatalf("GetClient() error = %v", err)
			}
			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
	}
}
