package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
)

// This is the main function of the program
func main() {

	if len(os.Args) > 2 {
		fmt.Println("error: either 0 or 1 arguments expected")
		return
	}
	port, portErr := findOpenPort(9000, 9100)
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

	peerMap := PeerMap{peers: make(map[string]Peer)}
	myPeer := Peer{addr: myAddr}
	peerMap.AddPeer(myPeer)

	go RequestHandler(myAddr, &peerMap)
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
			SendAddMeRequest(myAddr, seedAddr, &peerMap)
			SendMoreAddMeRequests(myAddr, seedAddr, &peerMap)
		}
	}

	exitChannel := make(chan bool)
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt)
	go func() {
		for range osSignals {
			exitChannel <- true
		}
	}()

	go RunHttpServer(&peerMap, exitChannel)
	go RunCommandParser(myAddr, &peerMap, exitChannel)

	for {
		if <-exitChannel {
			break
		}
	}

	peerMap.mutex.RLock()
	for _, peer := range peerMap.peers {
		if peer.addr != myAddr {
			SendRemoveMeRequest(myAddr, peer.addr)
		}
	}
	peerMap.mutex.RUnlock()
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

// SendAddMeRequest This function sends the AddMeRequest message to the seed address
func SendAddMeRequest(from net.Addr, to net.Addr, peerMap *PeerMap) {
	fmt.Println("connecting to:", to)
	conn, connErr := MakeTcpConnection(to)
	if connErr != nil {
		return
	}

	message := Message{Code: AddMeRequest, SenderAddr: from.String()}
	err := SendMessage(conn, message)
	if err != nil {
		return
	}

	err = conn.Close()
	peerMap.AddPeer(Peer{addr: to})
}

func SendRemoveMeRequest(from net.Addr, to net.Addr) {
	fmt.Println("disconnecting  from:", to)
	conn, connErr := MakeTcpConnection(to)
	if connErr != nil {
		return
	}

	message := Message{Code: RemoveMeRequest, SenderAddr: from.String()}
	err := SendMessage(conn, message)
	if err != nil {
		return
	}

	err = conn.Close()
}

func SendMoreAddMeRequests(from net.Addr, toPeersOf net.Addr, peerMap *PeerMap) {
	conn, connErr := MakeTcpConnection(toPeersOf)
	if connErr != nil {
		return
	}

	message := Message{Code: SharePeersRequest, SenderAddr: from.String()}
	err := SendMessage(conn, message)
	if err != nil {
		return
	}

	resp, respErr := ReceiveMessage(conn)

	if respErr != nil {
		return
	}

	for _, addrStr := range resp.Peers {
		if !peerMap.HasPeer(addrStr) {

			addr, addrParseErr := net.ResolveTCPAddr("tcp", addrStr)

			if addrParseErr != nil {
				fmt.Println("addr parse error:", addrParseErr)
				continue
			}

			SendAddMeRequest(from, addr, peerMap)
		}
	}
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

			err := SendMessage(conn, Message{SenderAddr: address.String()})

			if err != nil {
				fmt.Println("error sending address:", err)
				return
			}
		}
	}
}

func MakeTcpConnection(to net.Addr) (net.Conn, error) {
	duration, parseErr := time.ParseDuration("5s")

	if parseErr != nil {
		fmt.Println("duration parse error:", parseErr)
		return nil, parseErr
	}

	conn, connErr := net.DialTimeout("tcp", to.String(), duration)

	if connErr != nil {
		fmt.Println("error connecting to address:", connErr)
		return nil, connErr
	}

	return conn, nil
}

func SendBroadcastMessage(to net.Addr, from net.Addr, messageText string) {

	conn, connErr := MakeTcpConnection(to)

	if connErr != nil {
		return
	}

	var message = Message{
		Code:             BroadcastMessage,
		SenderAddr:       from.String(),
		BroadcastMessage: messageText,
	}

	err := SendMessage(conn, message)

	if err != nil {
		fmt.Println("error sending broadcast message:", err)
	}

	closeErr := conn.Close()

	if closeErr != nil {
		fmt.Println("error closing connection:", closeErr)
	}
}

// This is basically the go equivalent of an enum (a bunch of related constants)
const (
	PingRequest = iota
	PingResponse
	AddMeRequest
	SharePeersRequest
	SharePeersResponse
	BroadcastMessage
	RemoveMeRequest
)
