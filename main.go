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
	port, portErr := findOpenPort(8080, 8100)
	if portErr != nil {
		fmt.Println("error finding open port:", portErr)
		return
	}
	address := GetLocalIPAddress() + port
	myAddr, myAddrParseErr := net.ResolveTCPAddr("tcp", address)

	if myAddrParseErr != nil {
		fmt.Println("error parsing my ip arg:", myAddrParseErr)
		return
	}
	addresses := make(map[net.Addr]int)
	addrChan := make(chan net.Addr)
	go RequestHandler(myAddr, addrChan, addresses)
	go StoreAddresses(addresses, addrChan)
	fmt.Println("my address:", myAddr)
	//addrChan <- myAddr

	// If an address is given, try to join its network
	if len(os.Args) > 1 {
		seedAddrArg := os.Args[1]

		seedAddr, addrParseErr := net.ResolveTCPAddr("tcp", seedAddrArg)

		if addrParseErr != nil {
			fmt.Println("seedAddr parse error:", addrParseErr)
			return
		} else {
			SendAddMeRequest(addrChan, myAddr, seedAddr, addresses)
		}
	}

	// Runtime loop
	go ParseCommands(addrChan, myAddr, addresses)
	for {
	}
}

// GetLocalIPAddress /* This function returns the local IP address of the machine
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

// This function sends the AddMeRequest message to the seed address
func SendAddMeRequest(addrChan chan<- net.Addr, from net.Addr, to net.Addr, addresses map[net.Addr]int) {
	// Connect to the seed address
	fmt.Println("connecting to:", to)
	conn, connErr := net.Dial("tcp", to.String())

	if connErr != nil {
		fmt.Println("error connecting to seed address:", connErr)
		return
	}
	stringAddrs := NetAddrMapToStringMap(addresses)
	var message = Message{Code: AddMeRequest, SenderAddr: from.String(), ConnectedParties: stringAddrs}
	err := SendMessage(conn, message)

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

	connectedParties := StringMapToNetAddrMap(message.ConnectedParties)

	for eachAddr, _ := range connectedParties {
		addrChan <- eachAddr
	}

	addrChan <- messageAddr

	ShareAddress(from, connectedParties)
}

// ShareAddress /* This function sends the given address to all addresses in the given map
func ShareAddress(address net.Addr, addresses map[net.Addr]int) {
	for addr, _ := range addresses {
		if addr.String() != address.String() {
			conn, connErr := net.Dial("tcp", addr.String())

			if connErr != nil {
				fmt.Println("error connecting to address:", connErr)
				return
			}

			err := SendMessage(conn, Message{Code: ShareAddressRequest, SenderAddr: address.String(), ConnectedParties: NetAddrMapToStringMap(addresses)})

			if err != nil {
				fmt.Println("error sending address:", err)
				return
			}

		}
	}
}

// This is basically the go equivalent of an enum (a bunch of related constants)
const (
	PingRequest = iota
	PingResponse
	AddMeRequest
	AddMeResponse
	ShareAddressRequest
)
