package main

import (
	//lint:ignore ST1001 import bpsock
	. "bpsock-go/bpsock"
	//lint:ignore ST1001 import handler
	. "bpsock-go/handler"
	//lint:ignore ST1001 import tags
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

		// Create print Handler
		tagPrint := NewTag16("print")
		actionPrint := func(h Handler, tagName string, i int) {
			fmt.Println("Action Print")
			s := string(h.Data()[i])
			fmt.Println("data: ", s)
		}
		print := NewHookHandler(tagPrint, actionPrint)
		bpsock.AddHandler(&print)

		// Create Login Handler
		tagLogin := NewTag8("Login")
		actionLogin := func(h Handler, tagName string, i int) {
			fmt.Println("process login")
			fmt.Println("dataLogin: ", h.Data()[i])
			s := string(h.Data()[i])

			bpsock.SendResp([]byte("login ok"+s), tagName)

		}
		login := NewReqHandler(tagLogin, actionLogin)
		bpsock.AddHandler(&login)

	}
}
