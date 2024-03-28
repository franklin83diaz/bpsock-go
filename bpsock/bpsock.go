package bpsock

import (
	"bpsock-go/tags"
	"net"
)

// Bpsock
type BpSock struct {
	socket net.Conn
	dmtu   int
}

func (bpsock *BpSock) New(socket net.Conn, dmtu ...int) {

	bpsock.dmtu = 15000000
	if len(dmtu) > 0 {
		bpsock.dmtu = dmtu[0]
	}
	bpsock.socket = socket

}

func (bpsock *BpSock) send(data []byte, tag tags.Tag16) error {
	// !TODO:
	//lock up channel if it is busy
	//icrement channel counter

	//sendData(data, tag.name, _idChannel, bpsock.socket, bpsock.dmtu);

	return nil
}
