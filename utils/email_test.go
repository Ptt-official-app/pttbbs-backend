package utils

import "testing"

func Test_encodeRFC2047(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		str  string
		want string
	}{
		// TODO: Add test cases.
		{
			str:  "測試",
			want: "=?utf-8?b?5ris6Kmm?=",
		},
		{
			str:  "test",
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := encodeRFC2047(tt.str)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("encodeRFC2047() = %v, want %v", got, tt.want)
			}
		})
	}
}
