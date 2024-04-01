package bpsock

import (
	. "bpsock-go/handler"
	. "bpsock-go/tags"
	. "bpsock-go/utils"
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

// / Send data
// data: the data to send
// tag: the tag to use
// return: error
//
// use
// ```
// go func() {
// ch <- bpsock.Send(data, tag)
// }()
// ```
// to send data
func (bpsock *BpSock) Send(data []byte, tag Tag16) error {

	//reset channel if it is 65535
	if bpsock.id_chan == 65535 {
		bpsock.id_chan = 0
	}
	//icrement channel count
	mutex.Lock()
	bpsock.id_chan++
	mutex.Unlock()

	return SendData(data, tag, bpsock.id_chan, bpsock.socket, bpsock.dmtu)
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
			return fmt.Errorf("The tag %s is already in the list", handler.Tag().Name())
		}
	}

	bpsock.handlers = append(bpsock.handlers, handler)
	return nil
}

// Get list of handlers
func (bpsock *BpSock) GetHandlers() []Handler {
	return bpsock.handlers
}

func (bpsock *BpSock) received() {

	buffer := make([]byte, 1024)
	for {
		// Read data
		bytesRead, err := bpsock.socket.Read(buffer)
		if err != nil {
			fmt.Println("Error reading data: ", err)
			break
		}
		b := buffer[:bytesRead]

		idBytes := b[0:2]
		idChan := int(idBytes[0])<<8 | int(idBytes[1])
		//	fmt.Println("ID Chan: ", idChan)

		tagBytes := b[2:18]
		tagName := BytesToStringTrimNull(tagBytes)

		//	fmt.Println("Tag Name: ", tagName)

		sizeDataBytes := b[18:22]
		sizeData := int(sizeDataBytes[0])<<24 | int(sizeDataBytes[1])<<16 | int(sizeDataBytes[2])<<8 | int(sizeDataBytes[3])
		//	fmt.Println("Size Data: ", sizeData)

		data := b[22 : sizeData+22]
		//fmt.Println("Data: ", data)
		//fmt.Println("Data string: ", string(data))

		listHandlers := bpsock.handlers
		for i := 0; i < len(listHandlers); i++ {

			if listHandlers[i].Tag().Name() == tagName {

				action := listHandlers[i].ActionFunc()
				action(listHandlers[i], string(data), idChan)

			}
		}

	}
}
