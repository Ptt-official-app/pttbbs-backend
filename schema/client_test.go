package schema

import (
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
