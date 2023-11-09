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
			fmt.Println("connecting to:", seedAddr)
			conn, connErr := net.Dial("tcp", seedAddr.String())

			if connErr != nil {
				fmt.Println("error connecting to seed address:", connErr)
				return
			} else {
				_, err := fmt.Fprintf(conn, "%d\n%s\n", JoinRequest, myAddr.String())

				if err != nil {
					fmt.Println("error sending join request:", err)
					return
				}

				addrChan <- seedAddr
			}
		}

	}

	for {
	}
}

// This is basically the go equivalent of an enum (a bunch of related constants)
const (
	PingRequest = iota
	PingResponse
	JoinRequest
	JoinResponse
)
