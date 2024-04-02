package main

import (
	//. "bpsock-go/bpsock"
	. "bpsock-go/bpsock"
	. "bpsock-go/handler"
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

	for {

		// Accept a connection
		socket, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err)
			return
		}
		defer socket.Close()

		// Create a new BpSock object
		bpsock := NewBpSock(socket, 100)

		// Create a new tag
		tagLogin := NewTag16("Login")

		var actionLogin ActionFunc

		actionLogin = func(h Handler, tagName string, i int) {
			fmt.Println("Action Login")
			//fmt.Println("tag: ", tagName)
			s := string(h.Data()[i])
			fmt.Println(len(h.Data()))
			fmt.Println("data: ", s)
		}

		login := NewHookHandler(tagLogin, actionLogin)

		bpsock.AddHandler(&login)

		// buffer := make([]byte, 1024)
		// for {
		// 	// Read data
		// 	bytesRead, err := socket.Read(buffer)
		// 	if err != nil {
		// 		fmt.Println("Error reading data: ", err)
		// 		break
		// 	}
		// 	b := buffer[:bytesRead]
		// 	for i := 0; i < len(b); i++ {
		// 		fmt.Printf("%08b\n", b[i])
		// 	}
		// 	fmt.Println()
		// 	fmt.Println("Received data: ", buffer[:bytesRead])
		// }

	}
}
