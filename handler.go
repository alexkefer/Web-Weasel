package main

import (
	"fmt"
	"net"
)

func RequestHandler(myAddr net.Addr, addressChan chan<- net.Addr) {
	listener, listenErr := net.Listen("tcp", myAddr.String())

	if listenErr != nil {
		fmt.Println("tcp listen error:", listenErr)
		return
	}

	// This is an infinite loop in go
	for {
		conn, connErr := listener.Accept()

		if connErr != nil {
			fmt.Println("tcp connection error:", connErr)
		} else {
			// Here we create a separate goroutine (thread) to handle this connection
			go HandleConnection(myAddr, conn, addressChan)
		}
	}
}

func HandleConnection(myAddr net.Addr, conn net.Conn, addressChan chan<- net.Addr) {
	fmt.Println("connection received from:", conn.RemoteAddr())
	message, messageErr := ReceiveMessage(conn)

	if messageErr != nil {
		fmt.Println("message receive error:", messageErr)
		return
	}

	fmt.Printf("code: %d\n", message.Code)

	switch message.Code {
	case AddMeRequest:

		addrStr := message.SenderAddr
		fmt.Println("addr str:", addrStr)

		addr, addrParseErr := net.ResolveTCPAddr("tcp", addrStr)

		if addrParseErr != nil {
			fmt.Println("addr parse error:", addrParseErr)
			return
		}

		// Maybe ping addr here to make sure the address is legit

		addressChan <- addr

		messageErr := SendMessage(conn, Message{Code: AddMeResponse, SenderAddr: myAddr.String()})

		if messageErr != nil {
			fmt.Println("send AddMeResponse error:", messageErr)
			return
		}

	default:
		fmt.Printf("invalid code %d, closing connection.\n", message.Code)

	}

	CloseConnection(conn)
}

func CloseConnection(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		fmt.Println("connection close error:", err)
	} else {
		fmt.Println("closed connection:", conn.RemoteAddr())
	}
}
