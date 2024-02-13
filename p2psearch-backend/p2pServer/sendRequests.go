package p2pServer

import (
	"fmt"
	"github.com/alexkefer/p2psearch-backend/log"
	"github.com/alexkefer/p2psearch-backend/utils"
	"net"
)

// SendAddMeRequest This function sends the AddMeRequest message to the seed address
func SendAddMeRequest(from net.Addr, to net.Addr, peerMap *PeerMap) error {
	log.Info("connecting to seed address %s", to)
	conn, connErr := utils.MakeTcpConnection(to)
	if connErr != nil {
		return connErr
	}

	message := Message{Code: AddMeRequest, SenderAddr: from.String()}
	err := SendMessage(conn, message)
	if err != nil {
		return err
	}

	err = conn.Close()
	peerMap.AddPeer(Peer{Addr: to})

	return nil
}

func SendMoreAddMeRequests(from net.Addr, toPeersOf net.Addr, peerMap *PeerMap) {
	conn, connErr := utils.MakeTcpConnection(toPeersOf)
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
				log.Error("addr parse error: %s", addrParseErr)
				continue
			}

			SendAddMeRequest(from, addr, peerMap)
		}
	}
}

// Disconnect Sends RemoveMe requests to all peers
func Disconnect(myAddr net.Addr, peerMap *PeerMap) {
	log.Info("disconnecting from all peers")
	peerMap.Mutex.RLock()
	for _, peer := range peerMap.Peers {
		if peer.Addr != myAddr {
			sendRemoveMeRequest(myAddr, peer.Addr)
		}
	}
	peerMap.Mutex.RUnlock()
}

func sendRemoveMeRequest(from net.Addr, to net.Addr) {
	log.Info("disconnecting from: %s", to)
	conn, connErr := utils.MakeTcpConnection(to)
	if connErr != nil {
		return
	}

	message := Message{Code: RemoveMeRequest, SenderAddr: from.String()}
	err := SendMessage(conn, message)
	if err != nil {
		return
	}

	err = conn.Close()

	if err != nil {
		log.Error("failed to close connection")
	}
}

// ShareAddress /* This function sends the given address to all addresses in the given map
func ShareAddress(address net.Addr, addresses map[net.Addr]int) {
	for addr, _ := range addresses {
		if addr.String() != address.String() {
			conn, connErr := net.Dial("tcp", addr.String())

			if connErr != nil {
				log.Error("error connecting to address: %s", connErr)
				return
			}

			err := SendMessage(conn, Message{SenderAddr: address.String()})

			if err != nil {
				log.Error("error sending address: %s", err)
				return
			}
		}
	}
}

func SendBroadcastMessage(to net.Addr, from net.Addr, messageText string) {

	conn, connErr := utils.MakeTcpConnection(to)

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
