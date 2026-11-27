package dbcs

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const testForwardInfix = "\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc"

func TestMatchComment_forward(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    int
	}{
		{"forward at start", "\xa1\xb0 \x1b[1;32mabcd" + testForwardInfix + "\xac\xdd\xaaO Mavericks", 0},
		{"forward after newline", "x\n\xa1\xb0 \x1b[1;32mabcd" + testForwardInfix, 2},
		// the first marker is not a forward comment, the second one is: the offset must be relative to the whole content.
		{"offset after skipped marker", "abc def" + testForwardInfix + "\n\xa1\xb0 \x1b[1;32mabcd" + testForwardInfix, 7 + len(testForwardInfix) + 1},
		// the second "※ " is not at the beginning of a line: no match (and no infinite loop in ParseComments).
		{"skipped marker only", "abc def" + testForwardInfix + "\xa1\xb0 x" + testForwardInfix, -1},
		{"marker too close to the beginning", "ab" + testForwardInfix, -1},
		{"marker without leading space", "abc" + testForwardInfix, -1},
		{"no marker", "abc", -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MatchComment([]byte(tt.content)))
		})
	}
}

func TestMatchComment_deleted(t *testing.T) {
	deleted := "\x1b[1;30m(abc \xa7R\xb0\xa3 abc \xaa\xba\xb1\xc0\xa4\xe5: x)\x1b[m"
	tests := []struct {
		name    string
		content string
		want    int
	}{
		{"deleted at start", deleted, 0},
		{"deleted after newline", "x\n" + deleted, 2},
		// the first prefix is not a deleted comment: the offset must be relative to the whole content.
		{"offset after skipped prefix", "\x1b[1;30m(abc def\n" + deleted, len("\x1b[1;30m(abc def\n")},
		{"deleted not at the beginning of a line", "x " + deleted, -1},
		// a mid-line occurrence must not hide the comment on the next line.
		{"deleted after a mid-line occurrence", "x " + deleted + "\n" + deleted, len("x " + deleted + "\n")},
		{"no deleted", "abc", -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MatchComment([]byte(tt.content)))
		})
	}
}

// TestParseComments_noProgress: the forward marker appears twice in the same comment region.
// ParseComments used to loop forever (and allocate without bound) on such input.
func TestParseComments_noProgress(t *testing.T) {
	setupTest()
	defer teardownTest()

	commentsDBCS := []byte("x\nabc def" + testForwardInfix + "\xa1\xb0 x" + testForwardInfix)

	done := make(chan int, 1)
	go func() {
		done <- len(ParseComments("SYSOP", commentsDBCS, commentsDBCS))
	}()

	select {
	case n := <-done:
		assert.LessOrEqual(t, n, len(commentsDBCS))
	case <-time.After(5 * time.Second):
		t.Fatal("ParseComments does not terminate")
	}
}

func TestMatchComment_shortContent(t *testing.T) {
	assert.NotPanics(t, func() {
		MatchComment([]byte("abc" + testForwardInfix))
	})
}
