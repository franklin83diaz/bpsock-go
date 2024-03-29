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

	// Create a new BpSock object
	bpsock := NewBpSock(socket)

	// Create a new tag
	tag := NewTag16("Login")

	fmt.Println(bpsock)
	fmt.Println(tag)

}