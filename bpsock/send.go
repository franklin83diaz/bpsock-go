package bpsock

import (
	//lint:ignore ST1001 import tags
	. "bpsock-go/tags"
	"fmt"
)

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

func (bpsock *BpSock) SendResp(data []byte, tagName string) error {

	if len(tagName) > 15 {
		return fmt.Errorf("the tag '%s' is greater than 15 characters", tagName)
	}

	//reset channel if it is 65535
	if bpsock.id_chan == 65535 {
		bpsock.id_chan = 0
	}
	//icrement channel count
	mutex.Lock()
	bpsock.id_chan++
	mutex.Unlock()

	tagResp := NewTag16(tagName)

	// Send response type 2
	return SendData(data, tagResp, bpsock.id_chan, bpsock.socket, bpsock.dmtu, 2)
}
