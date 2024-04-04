package handler_test

import (
	. "bpsock-go/handler"
	. "bpsock-go/tags"
	"fmt"
	"testing"
)

func TestHookHandler_New(t *testing.T) {
	//actionFunc
	actionFunc := func(a Handler, b string, c int) {
		fmt.Println("actionFunc")
	}

	hookHandler3 := NewHookHandler(NewTag16("tag3"), nil)
	hookHandler4 := NewHookHandler(NewTag16("tag4"), actionFunc)
	var handler Handler = &hookHandler3

	if hookHandler3.Tag().Name() != "tag3" {
		t.Errorf("Expected tag to be 'tag3', got %s", hookHandler3.Tag())
	}

	if hookHandler4.Tag().Name() != "tag4" {
		t.Errorf("Expected tag to be 'tag4', got %s", hookHandler4.Tag())
	}

	if handler.ActionFunc() != nil {
		t.Errorf("Expected actionFunc to be nil, got %v", hookHandler3.ActionFunc())

	}

	//reqHandler
	reqHandler := NewReqHandler(NewTag8("tag5"), actionFunc)
	var handler2 Handler = &reqHandler

	if handler2.Tag().Name()[8:] != "tag5" {
		t.Errorf("Expected tag to be 'tag5', got %s", reqHandler.Tag().Name()[8:])
	}

}

func TestNewReqHandler(t *testing.T) {

	tests := []struct {
		name string
		args string
		want int
	}{
		// test cases.
		{"TestNewReqHandler", "tag5", 12},
		{"TestNewReqHandler", "tag12345", 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag8 := NewTag8(tt.args)
			if got := NewReqHandler(tag8, nil); len(got.Tag().Name()) != tt.want {
				t.Errorf("NewReqHandler() = %v, want %v", got, tt.want)
			}

		})
	}
}
