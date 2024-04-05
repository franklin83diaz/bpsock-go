package bpsock

import (
	//lint:ignore ST1001 import handler
	. "bpsock-go/handler"

	//lint:ignore ST1001 import tags
	. "bpsock-go/tags"

	"fmt"
)

// Request
func (bpsock *BpSock) Req(tag Tag8, data []byte, actionFunc ActionFunc) {
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

}

// Cancel the reqHandler
func (bpsock *BpSock) CancelReq(tag Tag16, id int) error {
	//get the handlers
	listHandlers := bpsock.handlers

	// Send cancel request
	err := SendData(nil, tag, id, bpsock.socket, bpsock.dmtu)
	if err != nil {
		return err
	}

	// Check if the tag is in the list of handlers
	for i := 0; i < len(listHandlers); i++ {

		// Remove Tag if is in the list of handlers
		if listHandlers[i].Tag().Name() == tag.Name() {
			//remove data from the handler after the action is executed
			defer listHandlers[i].RemoveData(id)
			return nil
		}
	}

	return fmt.Errorf("the tag %s is not in the list", tag.Name())
}
