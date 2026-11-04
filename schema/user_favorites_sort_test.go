package schema

import (
	"slices"
	"testing"
)

func TestSortUserFavoritesStable(t *testing.T) {
	first := &UserFavorites{FavIdx: 2}
	second := &UserFavorites{FavIdx: 1}
	third := &UserFavorites{FavIdx: 2}
	for _, ascending := range []bool{true, false} {
		input := []*UserFavorites{first, second, third}
		want := []*UserFavorites{second, first, third}
		if !ascending {
			want = []*UserFavorites{first, third, second}
		}
		got := SortUserFavoritesByFavIdx(input, ascending)
		if !slices.Equal(got, want) || !slices.Equal(input, want) {
			t.Errorf("ascending=%v: sort must preserve ties and update input in place", ascending)
		}
	}
	if got := SortUserFavoritesByFavIdx(nil, true); got != nil {
		t.Fatal("nil input must remain nil")
	}
	empty := []*UserFavorites{}
	if got := SortUserFavoritesByFavIdx(empty, false); got == nil || len(got) != 0 {
		t.Fatal("empty input must remain non-nil and empty")
	}
}
