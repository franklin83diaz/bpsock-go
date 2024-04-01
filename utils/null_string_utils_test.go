package utils_test

import (
	. "bpsock-go/utils"
	"testing"
)

func TestBytesToStringTrimNull(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Test string without null", args{[]byte("hello")}, "hello"},
		{"Test string with null", args{[]byte("hello\x00")}, "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToStringTrimNull(tt.args.b); got != tt.want {
				t.Errorf("BytesToStringTrimNull() = %v, want %v", got, tt.want)
			}
		})
	}
}
