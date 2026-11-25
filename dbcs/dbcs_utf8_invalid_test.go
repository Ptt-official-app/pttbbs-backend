package dbcs

import (
	"testing"

	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/stretchr/testify/assert"
)

// Test_dbcsToUtf8PerLine_invalidUtf8: an invalid UTF-8 byte right after a color code
// used to be measured as 3-byte U+FFFD, slicing past the end of the line (panic)
// or swallowing the following bytes.
func Test_dbcsToUtf8PerLine_invalidUtf8(t *testing.T) {
	setupTest()
	defer teardownTest()

	tests := []struct {
		name     string
		in       string
		wantDBCS []string
	}{
		{"color1 then invalid byte at end", "\x1b[131m\xb1", []string{"\x1b[131m\xb1"}},
		{"color1 then invalid byte then text", "\x1b[131m\xb1abc", []string{"\x1b[131m\xb1", "abc"}},
		{"color0 + color1 then invalid byte at end", "\x1b[33m\x1b[131m\xb1", []string{"\x1b[33m\x1b[131m\xb1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var line []*types.Rune
			assert.NotPanics(t, func() { line, _ = dbcsToUtf8PerLine(tt.in, types.DefaultColor) })
			got := make([]string, 0, len(line))
			for _, each := range line {
				got = append(got, each.DBCSStr)
			}
			assert.Equal(t, tt.wantDBCS, got)
		})
	}
}

func Test_dbcsToUtf8PerLineNextUtf8_invalidUtf8(t *testing.T) {
	assert.Equal(t, "\xb1", dbcsToUtf8PerLineNextUtf8("\xb1"))
	assert.Equal(t, "\xb1", dbcsToUtf8PerLineNextUtf8("\xb1abc"))
}
