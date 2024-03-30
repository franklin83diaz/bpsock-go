package bpsock

import (
	. "bpsock-go/handler"
	. "bpsock-go/tags"
	"fmt"
	"net"
)

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

	return &BpSock{
		socket: socket,
		dmtu:   defaultDmtu,
	}

}

func (bpsock *BpSock) Send(data []byte, tag Tag16) error {

	//icrement channel count
	bpsock.id_chan++

	//TODO: put send data in a goroutine

	SendData(data, tag, bpsock.id_chan, bpsock.socket, bpsock.dmtu)

	return nil
}

func (bpsock *BpSock) AddHandler(handler Handler) {
	bpsock.handlers = append(bpsock.handlers, handler)
}

func (bpsock *BpSock) Received() {

	buffer := make([]byte, 1024)
	for {
		// Read data
		bytesRead, err := bpsock.socket.Read(buffer)
		if err != nil {
			fmt.Println("Error reading data: ", err)
			break
		}
		b := buffer[:bytesRead]

		lenBytes := b[0:2]
		letInt := int(lenBytes[0])<<8 | int(lenBytes[1])
		fmt.Println("lenInt: ", letInt)

	}
}
