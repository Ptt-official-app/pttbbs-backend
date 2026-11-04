package api

import (
	"reflect"
	"sync"
	"testing"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func TestIndex(t *testing.T) {
	setupTest()
	defer teardownTest()

	params := &IndexParams{}

	type args struct {
		params any
	}
	tests := []struct {
		name     string
		args     args
		expected any
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{&IndexParams{}},
			expected: &IndexResult{Data: params},
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			user := &UserInfo{UserID: bbs.UUserID(pttbbsapi.GUEST), IsOver18: true}
			got, _, err := Index(testIP, user, tt.args.params, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("Index() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Index() = %v, want %v", got, tt.expected)
			}
		})
	}
	wg.Wait()
}
