package main

import (
	//lint:ignore ST1001 import bpsock
	. "bpsock-go/bpsock"
	"time"

	//lint:ignore ST1001 import tags
	. "bpsock-go/tags"
	//lint:ignore ST1001 import handler
	. "bpsock-go/handler"
	"fmt"
	"net"
)

func main() {

	//connect to the server
	//local testing
	socket, err := net.Dial("tcp", "localhost:8080")

	//netlab testting
	//socket, err := net.Dial("tcp", "192.168.137.254:8080")
	if err != nil {
		fmt.Println("Error connecting to server: ", err)
		return
	}
	defer socket.Close()

	//Create a new BpSock object
	bpsock := NewBpSock(socket, 100)

	//Create a new tag
	//tag := NewTag16("print")

	//bpsock.Send([]byte("hola hh1234567890 hola esto es un prueba de un string de mas de 100 runner espero haber escrito lo suficiente por segunda vez hola hh1234567890 hola esto es un prueba de un string de mas de 100 runner espero haber escrito lo suficiente listo."), tag)

	time.Sleep(1 * time.Second)

	//send request
	tag, id := bpsock.Req(NewTag8("Login"), []byte(`{"login": "pedro"}`), func(h Handler, tagName string, i int) {
		fmt.Println(" Login OK ")
		//remove handler after the action is executed
		bpsock.RemoveHandler(tagName)
	})
	time.Sleep(3 * time.Second)
	fmt.Println(tag.Name(), id)
	//bpsock.CancelReq(tag, id)
	time.Sleep(10 * time.Second)
}
