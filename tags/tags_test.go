package tags

import (
	"reflect"
	"testing"
)

func TestNewTag16(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want Tag16
	}{
		// test cases.
		{"TestNewTag16", args{"TestNewTag16"}, Tag16{"TestNewTag16"}},
		{"TestNewTag16", args{"1234567890123456"}, Tag16{"1234567890123456"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := NewTag16(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTag16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTag8(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want Tag8
	}{
		// test cases.
		{"TestNewTag8", args{"NewTag8"}, Tag8{"NewTag8"}},
		{"TestNewTag8", args{"12345678"}, Tag8{"12345678"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTag8(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTag8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTag16Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	NewTag16("12345678901234567")

}

func TestNewTag8Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	NewTag8("123456789")

}
