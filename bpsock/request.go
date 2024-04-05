package bpsock

import (
	//lint:ignore ST1001 import handler
	. "bpsock-go/handler"

	//lint:ignore ST1001 import tags
	. "bpsock-go/tags"

	"fmt"
)

// Request
// This function sends a request to the server
// Returns the tag and the id of the channel
func (bpsock *BpSock) Req(tag Tag8, data []byte, actionFunc ActionFunc) (Tag16, int) {
	// Create a new handler
	handler := NewReqHandler(tag, actionFunc)
	//add the handler to the list
	bpsock.AddHandler(Handler(&handler))

	// Increment channel count
	mutex.Lock()
	bpsock.id_chan++
	mutex.Unlock()

	// Send request type 1
	SendData(data, handler.Tag(), bpsock.id_chan, bpsock.socket, bpsock.dmtu, 1)

	return handler.Tag(), bpsock.id_chan
}

// Cancel the reqHandler
func (bpsock *BpSock) CancelReq(tag Tag16, id_chan int) error {
	//get the handlers
	listHandlers := bpsock.handlers

	// Send cancel request type 3
	err := SendData(nil, tag, id_chan, bpsock.socket, bpsock.dmtu, 3)
	if err != nil {
		return err
	}

	//remove handler
	for i := 0; i < len(listHandlers); i++ {
		if listHandlers[i].Tag().Name() == tag.Name() {
			bpsock.RemoveHandler(tag.Name())
			return nil
		}
	}

	return fmt.Errorf("handler not found")
}
