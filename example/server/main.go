package main

import (
	//. "bpsock-go/bpsock"
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
	defer socket.Close()

	// Create a new BpSock object
	//bpsock := NewBpSock(socket)

	// Create a new tag
	tag := NewTag16("Login")

	//fmt.Println(bpsock)
	fmt.Println(tag)

	buffer := make([]byte, 1024)
	for {
		// Read data
		bytesRead, err := socket.Read(buffer)
		if err != nil {
			fmt.Println("Error reading data: ", err)
			break
		}
		b := buffer[:bytesRead]
		for i := 0; i < len(b); i++ {
			fmt.Printf("%08b\n", b[i])
		}
		fmt.Println()
		fmt.Println("Received data: ", buffer[:bytesRead])
	}

}
