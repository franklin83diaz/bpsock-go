package handler_test

import (
	. "bpsock-go/handler"
	. "bpsock-go/tags"
	"fmt"
	"testing"
)

func TestHookHandler_New(t *testing.T) {
	//actionFunc
	//var actionFunc ActionFunc
	actionFunc := func(a Handler, b string, c int) {
		fmt.Println("actionFunc")
	}

	hookHandler3 := NewHookHandler(NewTag16("tag3"), nil)
	hookHandler4 := NewHookHandler(NewTag16("tag4"), actionFunc)

	if hookHandler3.Tag().Name() != "tag3" {
		t.Errorf("Expected tag to be 'tag3', got %s", hookHandler3.Tag())
	}

	if hookHandler4.Tag().Name() != "tag4" {
		t.Errorf("Expected tag to be 'tag4', got %s", hookHandler4.Tag())
	}

}
