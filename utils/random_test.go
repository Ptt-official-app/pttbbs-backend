package utils

import (
	"errors"
	"testing"
)

func TestGenRandomInt64(t *testing.T) {
	for i := 0; i < 100; i++ {
		got, err := GenRandomInt64(10)
		if err != nil {
			t.Fatalf("GenRandomInt64(10) e: %v", err)
		}
		if got < 0 || got >= 10 {
			t.Errorf("GenRandomInt64(10) = %v, out of range", got)
		}
	}

	for _, theMax := range []int64{0, -1} {
		if _, err := GenRandomInt64(theMax); !errors.Is(err, ErrInvalidMax) {
			t.Errorf("GenRandomInt64(%v) e: %v, want ErrInvalidMax", theMax, err)
		}
	}
}
