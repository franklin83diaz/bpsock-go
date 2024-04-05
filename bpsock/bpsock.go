package bpsock

import (
	//lint:ignore ST1001 import handler
	. "bpsock-go/handler"
	"fmt"
	"net"
	"sync"
)

var mutex sync.Mutex

// Bpsock
type BpSock struct {
	socket   net.Conn
	dmtu     int
	handlers []Handler
	id_chan  int
}

// Create a new BpSock object
// socket: the socket to use
// dmtu: the maximum size of the data to send
//
//	     (default is 15000000)
//		     dmtu max is 15000000
//
// return: a new BpSock object
func NewBpSock(socket net.Conn, dmtu ...int) *BpSock {

	// set the default dmtu
	defaultDmtu := 15000000
	if len(dmtu) > 0 {
		defaultDmtu = dmtu[0]
	}

	// check if the dmtu is greater than the maximum size
	if len(dmtu) > 15000000 {
		panic("the DMTU exceeds the maximum size of 15,000,000 bytes.")
	}

	bpSock := &BpSock{
		socket: socket,
		dmtu:   defaultDmtu,
	}

	go bpSock.received()
	return bpSock

}

// Close the BpSock
// this will close the socket
// and stop the received function
func (bpsock *BpSock) Close() {
	//Close the socket
	//when the socket is closed, the received function will stop
	bpsock.socket.Close()
}

// Add a handler to the BpSock object
func (bpsock *BpSock) AddHandler(handler Handler) error {
	//check if the handler tag is already in the list
	for i := 0; i < len(bpsock.handlers); i++ {
		if bpsock.handlers[i].Tag().Name() == handler.Tag().Name() {
			return fmt.Errorf("the tag %s is already in the list", handler.Tag().Name())
		}
	}

	bpsock.handlers = append(bpsock.handlers, handler)
	return nil
}

// Get list of handlers
func (bpsock *BpSock) GetHandlers() []Handler {
	return bpsock.handlers
}

// Remove a handler from the BpSock object
func (bpsock *BpSock) RemoveHandler(tagName string) error {
	//check if the handler tag is already in the list
	for i := 0; i < len(bpsock.handlers); i++ {
		if bpsock.handlers[i].Tag().Name() == tagName {
			bpsock.handlers = append(bpsock.handlers[:i], bpsock.handlers[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("the tag %s is not in the list", tagName)
}
