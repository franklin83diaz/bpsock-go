package main

import (
	. "bpsock-go/bpsock"
	. "bpsock-go/tags"

	"fmt"
	"net"
)

func main() {

	//create a new server
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error creating server: ", err)
		return
	}

	// Accept a connection
	socket, err := ln.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err)
		return
	}

	// Create a new BpSock object
	bpsock := NewBpSock(socket)

	// Create a new tag
	tag := NewTag16("Login")

	fmt.Println(bpsock)
	fmt.Println(tag)

}
