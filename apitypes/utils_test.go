package apitypes

import (
	"testing"

	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func TestToURL(t *testing.T) {
	type args struct {
		fbboardID  FBoardID
		farticleID FArticleID
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should be {URL}/board/Baseball/article/M.1629685889.A.AEB",
			args: args{fbboardID: "Baseball", farticleID: "M.1629685889.A.AEB"},
			want: types.FRONTEND_PREFIX + "/board/Baseball/article/M.1629685889.A.AEB",
		},
		{
			name: "null board should return empty string",
			args: args{fbboardID: "", farticleID: "M.1629685889.A.AEB"},
			want: "",
		},
		{
			name: "null article should return empty string",
			args: args{fbboardID: "Baseball", farticleID: ""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToURL(tt.args.fbboardID, tt.args.farticleID); got != tt.want {
				t.Errorf("ToURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
