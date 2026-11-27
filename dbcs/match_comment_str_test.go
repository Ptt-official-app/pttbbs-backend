package dbcs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchCommentStr(t *testing.T) {
	forward := "※ \x1b[1;32mabcd\x1b[0;32m:轉錄至看板 SYSOP\x1b[m"
	recommend := "\x1b[1;37m推 \x1b[33mabcd\x1b[m\x1b[33m: 讚\x1b[m"
	tests := []struct {
		name    string
		content string
		want    int
	}{
		{"forward at start", forward, 0},
		// the forward comment is not at the beginning of content, and it comes before the recommend.
		{"forward after lines of content", "line1\nline2\n" + forward + "\n" + recommend, len("line1\nline2\n")},
		{"forward not at the beginning of a line", "x " + forward, -1},
		{"no comment", "abc\ndef", -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MatchCommentStr(tt.content))
		})
	}
}
