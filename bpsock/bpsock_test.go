package bpsock_test

import (
	. "bpsock-go/bpsock"
	. "bpsock-go/handler"
	. "bpsock-go/tags"
	"net"
	"reflect"
	"testing"
)

// Test NewBpSock
func TestNewBpSock(t *testing.T) {
	type args struct {
		socket net.Conn
		dmtu   []int
	}
	socket := &net.TCPConn{}
	defer socket.Close()
	i := make([]int, 1)
	i[0] = 10000
	tests := []struct {
		name string
		args args
		want *BpSock
	}{
		// TODO: Add test cases.
		{"TestNewBpSock", args{nil, nil}, NewBpSock(nil, 15000000)},
		{"TestNewBpSock", args{socket, i}, NewBpSock(socket, 10000)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBpSock(tt.args.socket, tt.args.dmtu...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBpSock() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test add handler
func TestBpSock_AddHandler(t *testing.T) {

	socket := &net.TCPConn{}
	defer socket.Close()
	bpsock := NewBpSock(socket)
	tag := NewTag16("tagTest")
	handler := NewHookHandler(tag, (func(h Handler, s string, i int) {}))
	handlerTest := Handler(&handler)
	//2
	tag2 := NewTag16("tagTest2")
	handler2 := NewHookHandler(tag2, (func(h Handler, s string, i int) {}))
	handlerTest2 := &handler2

	tests := []struct {
		name  string
		args  Handler
		want  Handler
		want2 int // number of handlers
	}{
		{"TestAddHandler", handlerTest, handlerTest, 1},
		{"TestAddHandler same", handlerTest, handlerTest, 1},
		{"TestAddHandler 2", handlerTest2, handlerTest2, 2},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bpsock.AddHandler(tt.args)
			n := i
			if i > n-1 {
				n = len(bpsock.GetHandlers()) - 1
			}
			//test handler added
			if got := bpsock.GetHandlers()[n]; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddHandler() = %v, want %v", got, tt.want)
			}
			//test number of handlers
			if got := len(bpsock.GetHandlers()); got != tt.want2 {
				t.Errorf("AddHandler() = %v, want %v", got, tt.want2)
			}

		})
	}
}
