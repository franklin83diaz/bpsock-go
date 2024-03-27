package bpsock

import (
	"net"
)

// Tag16
type Tag16 struct {
	name string
}

func (t Tag16) New(s string) {
	if len(s) > 16 {
		panic("Tag16: tag name too long")
	}
	t.name = s
}

// Bpsock
type Bpsock struct {
	socket net.Conn
	dmtu   int
}

func (bpsock *Bpsock) New(socket net.Conn, dmtu ...int) {
	bpsock.dmtu = 15000000
	if len(dmtu) > 0 {
		bpsock.dmtu = dmtu[0]
	}
	bpsock.socket = socket
}

func (bpsock *Bpsock) send(data []byte, tag Tag16) error {
	// !TODO
	//lock up channel if it is busy
	//icrement channel counter

	//sendData(data, tag.name, _idChannel, bpsock.socket, bpsock.dmtu);

	return nil
}
