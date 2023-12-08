package main

import (
	"fmt"
	"net"
)

func RequestHandler(myAddr net.Addr, peerMap *PeerMap) {
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
			go HandleConnection(myAddr, conn, peerMap)
		}
	}
}

func HandleConnection(myAddr net.Addr, conn net.Conn, peerMap *PeerMap) {
	message, messageErr := ReceiveMessage(conn)

	if messageErr != nil {
		fmt.Println("message receive error:", messageErr)
		return
	}

	switch message.Code {

	case AddMeRequest:
		// Peer is asking us to add them to our peer map
		addrStr := message.SenderAddr

		peerAddr, addrParseErr := net.ResolveTCPAddr("tcp", addrStr)

		if addrParseErr != nil {
			fmt.Println("addr parse error:", addrParseErr)
			return
		}

		peer := Peer{addr: peerAddr}
		peerMap.AddPeer(peer)

	case SharePeersRequest:
		// Peer is asking us for our peer map

		addrStr := message.SenderAddr
		_, addrParseErr := net.ResolveTCPAddr("tcp", addrStr)

		if addrParseErr != nil {
			fmt.Println("addr parse error:", addrParseErr)
			CloseConnection(conn)
			return
		}

		peerSlice := make([]string, 0)

		for addr, _ := range peerMap.peers {
			peerSlice = append(peerSlice, addr)
		}

		peerMap.mutex.RLock()
		message = Message{Code: SharePeersResponse, Peers: peerSlice, SenderAddr: myAddr.String()}
		_ = SendMessage(conn, message)
		peerMap.mutex.RUnlock()

	case BroadcastMessage:
		fmt.Printf("received broadcast message from %s: %s\n", message.SenderAddr, message.BroadcastMessage)

	default:
		fmt.Printf("invalid code %d, closing connection.\n", message.Code)

	}

	CloseConnection(conn)
}

func CloseConnection(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		fmt.Println("connection close error:", err)
	}
}
