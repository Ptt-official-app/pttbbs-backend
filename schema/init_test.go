package schema

import (
	"reflect"
	"testing"
)

func Test_getBSONName(t *testing.T) {
	type args struct {
		empty     any
		fieldName string
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		// TODO: Add test cases.
		{
			args:     args{empty: EMPTY_ARTICLE, fieldName: "BBoardID"},
			expected: "bid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBSONName(tt.args.empty, tt.args.fieldName); got != tt.expected {
				t.Errorf("getBSONName() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func Test_getFields(t *testing.T) {
	expected := make(map[string]bool)
	expected[ARTICLE_CONTENT_MTIME_b] = true
	expected[ARTICLE_IS_DELETED_b] = true
	type args struct {
		empty    any
		fields_i any
	}
	tests := []struct {
		name           string
		args           args
		expectedFields map[string]bool
	}{
		// TODO: Add test cases.
		{
			args:           args{empty: EMPTY_ARTICLE, fields_i: EMPTY_ARTICLE_CONTENT_MTIME},
			expectedFields: expected,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFields := getFields(tt.args.empty, tt.args.fields_i); !reflect.DeepEqual(gotFields, tt.expectedFields) {
				t.Errorf("getFields() = %v, want %v", gotFields, tt.expectedFields)
			}
		})
	}
}

func Test_assertFields(t *testing.T) {
	type args struct {
		empty    any
		fields_i any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{EMPTY_ARTICLE, EMPTY_ARTICLE_QUERY},
		},
		{
			args: args{EMPTY_ARTICLE, EMPTY_ARTICLE_CONTENT_MTIME},
		},
		{
			args: args{EMPTY_ARTICLE, EMPTY_ARTICLE_CONTENT_INFO},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := assertFields(tt.args.empty, tt.args.fields_i); (err != nil) != tt.wantErr {
				t.Errorf("assertFields() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_assertAllFields(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := assertAllFields(); (err != nil) != tt.wantErr {
				t.Errorf("assertAllFields() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
