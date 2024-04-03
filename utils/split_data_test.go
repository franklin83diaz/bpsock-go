package utils_test

import (
	. "bpsock-go/utils"
	"reflect"
	"testing"
)

func TestSplitData(t *testing.T) {
	type args struct {
		data []byte
		dmtu int
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		// test cases.
		{
			name: "Test case 1",
			args: args{
				data: []byte("Hello World"),
				dmtu: 5,
			},
			want: [][]byte{
				[]byte("Hello"),
				[]byte(" Worl"),
				[]byte("d"),
			},
		},
		{
			name: "Test case 2",
			args: args{
				data: []byte("HelloWorld"),
				dmtu: 5,
			},
			want: [][]byte{
				[]byte("Hello"),
				[]byte("World"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitData(tt.args.data, tt.args.dmtu); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitData() = %v, want %v", got, tt.want)
			}
		})
	}
}
