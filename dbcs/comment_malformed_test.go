package dbcs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestParseComments_malformed: truncated / malformed comment lines must not panic.
func TestParseComments_malformed(t *testing.T) {
	setupTest()
	defer teardownTest()

	edit := "\xa1\xb0 \xbds\xbf\xe8: "
	forward := "\xa1\xb0 \x1b[1;32mabcd\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc"
	recommend := "\x1b[1;37m\xb1\xc0 \x1b[33mabcd\x1b[m\x1b[33m: 77"

	tests := []struct {
		name string
		in   string
	}{
		{"edit: empty", edit},
		{"edit: owner only", edit + "abcd "},
		{"edit: no closing paren", edit + "abcd (1.2.3.4"},
		{"edit: no date", edit + "abcd (1.2.3.4)"},
		{"edit: old format no closing paren", edit + "abcd \xa8\xd3\xa6\xdb: 1.2.3.4 (03/18"},
		{"reply then edit: empty user", "reply\n" + edit + "(1.2.3.4), 03/21/2021 03:04:47\n"},
		{"reply then edit: truncated date", "reply\n" + edit + "abcd (1.2.3.4), 03/21/2021"},
		{"forward: ends after board prefix", forward + "\xac\xdd\xaaO "},
		{"forward: ends after marker", forward},
		{"comment: ESC right after infix", recommend[:len(recommend)-3] + "\x1b"},
		{"comment: ESC directly after infix", "\x1b[1;37m\xb1\xc0 \x1b[33m\x1b[m\x1b[33m\x1b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := []byte(tt.in)
			assert.NotPanics(t, func() { ParseComments("SYSOP", in, in) })
			in2 := []byte(recommend + "\n" + tt.in)
			assert.NotPanics(t, func() { ParseComments("SYSOP", in2, in2) })
		})
	}
}
