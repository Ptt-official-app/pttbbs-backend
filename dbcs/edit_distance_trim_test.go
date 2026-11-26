package dbcs

import (
	"math/rand"
	"runtime"
	"strconv"
	"testing"

	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/stretchr/testify/assert"
)

// Test_calcEDInfos_sameAsFullDP: trimming the common prefix must not change the result of the full DP + backtracking.
func Test_calcEDInfos_sameAsFullDP(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	for range 20000 {
		nNew, nOrig := 1+r.Intn(7), 1+r.Intn(7)
		nAlphabet := 1 + r.Intn(3) // a small alphabet: many ties and repeated MD5s.
		newComments := make([]*schema.Comment, nNew)
		origComments := make([]*schema.CommentMD5, nOrig)
		for idx := range newComments {
			newComments[idx] = &schema.Comment{MD5: strconv.Itoa(r.Intn(nAlphabet))}
		}
		for idx := range origComments {
			origComments[idx] = &schema.CommentMD5{MD5: strconv.Itoa(r.Intn(nAlphabet)), SortTime: 1}
		}

		expected := calcEditDistanceBacktracking(calcEditDistance(newComments, origComments), newComments, origComments)
		got := calcEDInfos(newComments, origComments)

		if !assert.Equal(t, expected, got, "new: %v orig: %v", newComments, origComments) {
			return
		}
	}
}

// TestCalcEDBlocks_appendOnlyMemory: appending to a long comment list must not allocate an O(N*M) matrix.
func TestCalcEDBlocks_appendOnlyMemory(t *testing.T) {
	const nOrig = 5000

	newComments := make([]*schema.Comment, nOrig+1)
	origComments := make([]*schema.CommentMD5, nOrig)
	for idx := range origComments {
		md5 := strconv.Itoa(idx)
		newComments[idx] = &schema.Comment{MD5: md5, SortTime: 2}
		origComments[idx] = &schema.CommentMD5{MD5: md5, SortTime: 2}
	}
	newComments[nOrig] = &schema.Comment{MD5: "new", SortTime: 3}

	var before, after runtime.MemStats
	runtime.ReadMemStats(&before)
	edBlocks, err := CalcEDBlocks(newComments, origComments, 1, 4)
	runtime.ReadMemStats(&after)

	assert.NoError(t, err)
	assert.NotEmpty(t, edBlocks)
	// the full matrix would need about nOrig*nOrig*8 = 200MB.
	assert.Less(t, after.TotalAlloc-before.TotalAlloc, uint64(20<<20))
}
