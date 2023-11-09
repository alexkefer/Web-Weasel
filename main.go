package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("error: first arg must be own ip")
		return
	}

	myAddr, myAddrParseErr := net.ResolveTCPAddr("tcp", os.Args[1])

	if myAddrParseErr != nil {
		fmt.Println("error parsing my ip arg:", myAddrParseErr)
		return
	}

	addrChan := make(chan net.Addr)
	go RequestHandler(myAddr, addrChan)
	go StoreAddresses(addrChan)
	addrChan <- myAddr

	// If a second address is given, try to join its network
	if len(os.Args) > 2 {
		seedAddrArg := os.Args[2]

		seedAddr, addrParseErr := net.ResolveTCPAddr("tcp", seedAddrArg)

		if addrParseErr != nil {
			fmt.Println("seedAddr parse error:", addrParseErr)
			return
		} else {
			SendAddMeRequest(addrChan, myAddr, seedAddr)
		}
	}

	for {
	}
}

func SendAddMeRequest(addrChan chan<- net.Addr, from net.Addr, to net.Addr) {
	fmt.Println("connecting to:", to)
	conn, connErr := net.Dial("tcp", to.String())

	if connErr != nil {
		fmt.Println("error connecting to seed address:", connErr)
		return
	} else {
		err := SendMessage(conn, Message{Code: AddMeRequest, SenderAddr: from.String()})

		if err != nil {
			fmt.Println("error sending join request:", err)
			return
		}

		message, messageErr := ReceiveMessage(conn)

		if messageErr != nil {
			fmt.Println("error receiving add me response:", err)
			return
		}

		if message.Code != AddMeResponse {
			fmt.Println("unexpected message code:", message.Code)
			return
		}

		messageAddr, addrParseErr := net.ResolveTCPAddr("tcp", message.SenderAddr)

		if addrParseErr != nil {
			fmt.Println("addr parse error:", addrParseErr)
			return
		}

		addrChan <- messageAddr

	}
}

// This is basically the go equivalent of an enum (a bunch of related constants)
const (
	PingRequest = iota
	PingResponse
	AddMeRequest
	AddMeResponse
)
