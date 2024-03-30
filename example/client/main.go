package main

import (
	. "bpsock-go/bpsock"
	. "bpsock-go/tags"
	"fmt"
	"net"
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
	bpsock := NewBpSock(socket)

	//Create a new tag
	tag := NewTag16("Login")

	bpsock.Send([]byte("Hello, server!"), tag)

}
