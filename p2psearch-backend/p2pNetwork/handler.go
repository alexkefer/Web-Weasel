package p2pNetwork

import (
	"github.com/alexkefer/p2psearch-backend/log"
	"net"
)

func StartServer(myAddr net.Addr, peerMap *PeerMap) {
	listener, listenErr := net.Listen("tcp", myAddr.String())

	if listenErr != nil {
		log.Error("tcp listen error: %s", listenErr)
		return
	}

	// This is an infinite loop in go
	for {
		conn, connErr := listener.Accept()

		if connErr != nil {
			log.Error("tcp connection error: %s", connErr)
		} else {
			// Here we create a separate goroutine (thread) to handle this connection
			go HandleConnection(myAddr, conn, peerMap)
		}
	}
}

func HandleConnection(myAddr net.Addr, conn net.Conn, peerMap *PeerMap) {
	message, messageErr := ReceiveMessage(conn)

	if messageErr != nil {
		log.Error("message receive error: %s", messageErr)
		return
	}

	switch message.Code {

	case AddMeRequest:
		// Peer is asking us to add them to our p2pServer map
		addrStr := message.SenderAddr

		peerAddr, addrParseErr := net.ResolveTCPAddr("tcp", addrStr)

		if addrParseErr != nil {
			log.Error("addr parse error: %s", addrParseErr)
			return
		}

		peer := Peer{Addr: peerAddr}
		peerMap.AddPeer(peer)

	case SharePeersRequest:
		// Peer is asking us for our p2pServer map

		addrStr := message.SenderAddr
		_, addrParseErr := net.ResolveTCPAddr("tcp", addrStr)

		if addrParseErr != nil {
			log.Error("addr parse error: %s", addrParseErr)
			CloseConnection(conn)
			return
		}

		peerSlice := make([]string, 0)

		for addr, _ := range peerMap.Peers {
			peerSlice = append(peerSlice, addr)
		}

		peerMap.Mutex.RLock()
		message = Message{Code: SharePeersResponse, Peers: peerSlice, SenderAddr: myAddr.String()}
		_ = SendMessage(conn, message)
		peerMap.Mutex.RUnlock()

	case BroadcastMessage:
		log.Info("received broadcast message from %s: %s", message.SenderAddr, message.BroadcastMessage)

	case RemoveMeRequest:
		peerMap.RemovePeer(message.SenderAddr)

	default:
		log.Warn("invalid code %d, closing connection", message.Code)

	}

	CloseConnection(conn)
}

func CloseConnection(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		log.Error("connection close error: %s", err)
	}
}
