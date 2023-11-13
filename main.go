package main

import (
	"fmt"
	"net"
	"os"
)

// This is the main function of the program
func main() {

	if len(os.Args) > 2 {
		fmt.Println("error: either 0 or 1 arguments expected")
		return
	}
	port, err := findOpenPort(8080, 8100)
	if err != nil {
		fmt.Println("error finding open port:", err)
		return
	}
	address := GetLocalIPAddress() + port
	myAddr, myAddrParseErr := net.ResolveTCPAddr("tcp", address)

	if myAddrParseErr != nil {
		fmt.Println("error parsing my ip arg:", myAddrParseErr)
		return
	}

	addrChan := make(chan net.Addr)
	go RequestHandler(myAddr, addrChan)
	go StoreAddresses(addrChan)
	addrChan <- myAddr

	// If a second address is given, try to join its network
	if len(os.Args) > 1 {
		seedAddrArg := os.Args[1]

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

// This function returns the local IP address of the machine
// It is used to determine the address of the machine running the program so that other machines can connect to it
func GetLocalIPAddress() string {
	connection, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic("error getting local ip address")
	}
	defer connection.Close()
	localAddr := connection.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

// This function finds an open port in the range [startPort, endPort]
func findOpenPort(startPort, endPort int) (string, error) {
	for port := startPort; port <= endPort; port++ {
		// Attempt to bind to this port
		listener, listenerErr := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if listenerErr == nil {
			// If we were able to bind, close the listener and return the port
			listener.Close()
			return fmt.Sprintf(":%d", port), nil
		}
	}
	// If we were unable to bind to any ports, return an empty string
	return "", fmt.Errorf("unable to find open port")
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
