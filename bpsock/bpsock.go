package bpsock

import (
	. "bpsock-go/handler"
	. "bpsock-go/tags"
	"net"
)

// Bpsock
type BpSock struct {
	socket   net.Conn
	dmtu     int
	handlers []Handler
	id_chan  int
}

func NewBpSock(socket net.Conn, dmtu ...int) *BpSock {

	defaultDmtu := 15000000
	if len(dmtu) > 0 {
		defaultDmtu = dmtu[0]
	}

	return &BpSock{
		socket: socket,
		dmtu:   defaultDmtu,
	}

}

func (bpsock *BpSock) send(data []byte, tag Tag16) error {
	// !TODO:
	//lock up channel if it is busy
	//icrement channel counter

	//sendData(data, tag.name, _idChannel, bpsock.socket, bpsock.dmtu);

	return nil
}
