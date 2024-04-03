package bpsock

import (
	//lint:ignore ST1001 import handler
	. "bpsock-go/handler"
	//lint:ignore ST1001 import tags
	. "bpsock-go/tags"
	//lint:ignore ST1001 import utils
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

func (bpsock *BpSock) received() {

	start := 0
	end := 22
	buffer := make([]byte, bpsock.dmtu+22)
	var idChan int
	var tagName string
	var sizeData int
	for {
		// Read data
		bytesRead, err := bpsock.socket.Read(buffer[start:end])

		if err != nil {
			//if the error is EOF, then the socket is closed
			// no need to print the error
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading data: ", err)
			break
		}

		sizeUnit := bytesRead

		if start == 0 {
			b := buffer[:bytesRead]
			//id
			idBytes := b[0:2]
			idChan = int(idBytes[0])<<8 | int(idBytes[1])

			//tag
			tagBytes := b[2:18]
			tagName = BytesToStringTrimNull(tagBytes)

			//size data
			sizeDataBytes := b[18:22]
			sizeData = int(sizeDataBytes[0])<<24 | int(sizeDataBytes[1])<<16 | int(sizeDataBytes[2])<<8 | int(sizeDataBytes[3])

			//sizeUnit is the size of the data plus the header
			sizeUnit = sizeData + 22
		}

		//if the size of the data is greater than the bytes read
		if sizeUnit > bytesRead {
			start = bytesRead
			end = sizeData + 22
			continue
		}

		//reset start and end
		start = 0
		end = 22

		//if is end channel

		//data
		data := buffer[22 : sizeData+22]

		//get the handlers
		listHandlers := bpsock.handlers

		//check if the tag is in the list of handlers
		for i := 0; i < len(listHandlers); i++ {

			//if the tag is in the list of handlers
			if listHandlers[i].Tag().Name() == tagName {

				//if sizeData is 0, then it is an end channel
				if sizeData == 0 {
					action := listHandlers[i].ActionFunc()
					action(listHandlers[i], tagName, idChan)
					//just one handler per tag, no need to continue
					break
				}

				//add data to the handler
				listHandlers[i].AddData(idChan, data)

				//just one handler per tag, no continue to the next handler
				break

			}
		}

	}
}
