package main

import (
	"fmt"
	"net"
)

func RequestHandler(myAddr net.Addr, addressChan chan<- net.Addr, connectedParties map[net.Addr]int) {
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
			go HandleConnection(myAddr, conn, addressChan, connectedParties)
		}
	}
}

func HandleConnection(myAddr net.Addr, conn net.Conn, addressChan chan<- net.Addr, connectedParties map[net.Addr]int) {
	message, messageErr := ReceiveMessage(conn)

	if messageErr != nil {
		fmt.Println("message receive error:", messageErr)
		return
	}

	fmt.Printf("code: %d\n", message.Code)

	switch message.Code {
	case AddMeRequest:

		newConnections := StringMapToNetAddrMap(message.ConnectedParties)
		fmt.Printf("Connected Parties %v\n", message.ConnectedParties)

		addrStr := message.SenderAddr
		fmt.Println("addr str:", addrStr)

		addr, addrParseErr := net.ResolveTCPAddr("tcp", addrStr)

		if addrParseErr != nil {
			fmt.Println("addr parse error:", addrParseErr)
			return
		}
		fmt.Println("Connection from ", addr)

		// Maybe ping addr here to make sure the address is legit
		messageErr := SendMessage(conn, Message{Code: AddMeResponse, SenderAddr: myAddr.String(), ConnectedParties: NetAddrMapToStringMap(connectedParties)})

		if messageErr != nil {
			fmt.Println("send AddMeResponse error:", messageErr)
			return
		}

		for eachAddr, _ := range newConnections {
			if eachAddr.String() != myAddr.String() {
				fmt.Println("sending add me request to:", eachAddr)
				addressChan <- eachAddr
			}
		}

		addressChan <- addr
	case ShareAddressRequest:
		// Check if the address is already in the connected parties map and add it if it isn't
		addrStr := message.SenderAddr
		fmt.Println("addr str:", addrStr)

		addr, addrParseErr := net.ResolveTCPAddr("tcp", addrStr)

		if addrParseErr != nil {
			fmt.Println("addr parse error:", addrParseErr)
			return
		}

		_, ok := connectedParties[addr]

		if !ok {
			connectedParties[addr] = 0
			fmt.Printf("address added to store: %s, addr store size: %d\n", addr, len(connectedParties))
			fmt.Printf("Entire Network: %v\n", connectedParties)
		}

		if messageErr != nil {
			fmt.Println("send ShareAddressResponse error:", messageErr)
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
