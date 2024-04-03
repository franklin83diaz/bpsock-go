package main

import (
	. "bpsock-go/bpsock"
	. "bpsock-go/tags"
	"fmt"
	"net"
	"time"
)

func main() {

	//connect to the server
	socket, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server: ", err)
		return
	}
	defer socket.Close()

	//Create a new BpSock object
	bpsock := NewBpSock(socket, 100)

	//Create a new tag
	tag := NewTag16("Login")

	bpsock.Send([]byte("1234567890 hola esto es un prueba de un string de mas de 100 runner espero haber escrito lo suficiente"), tag)
	time.Sleep(1000 * time.Millisecond)

}
