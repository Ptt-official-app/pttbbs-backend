package configutil

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestSetListStringConfig(t *testing.T) {
	orig := myViper
	defer func() { myViper = orig }()

	tests := []struct {
		name  string
		value any // nil: key is not set
		orig  []string
		want  []string
	}{
		{name: "not set keeps orig", orig: []string{"*"}, want: []string{"*"}},
		{name: "empty value", value: "", orig: []string{"*"}, want: []string{}},
		{name: "single", value: "a", want: []string{"a"}},
		{name: "trims and drops empty items", value: " a , ,b,, ", want: []string{"a", "b"}},
		{name: "only separators", value: " , ", want: []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			myViper = viper.New()
			if tt.value != nil {
				myViper.Set("prefix.key", tt.value)
			}

			got := SetListStringConfig("prefix", "KEY", tt.orig)
			assert.Equal(t, tt.want, got)
		})
	}
}
