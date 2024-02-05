package p2pServer

import (
	"fmt"
	"github.com/alexkefer/p2psearch-backend/utils"
	"net"
)

// SendAddMeRequest This function sends the AddMeRequest message to the seed address
func SendAddMeRequest(from net.Addr, to net.Addr, peerMap *PeerMap) {
	fmt.Println("connecting to:", to)
	conn, connErr := utils.MakeTcpConnection(to)
	if connErr != nil {
		return
	}

	message := Message{Code: AddMeRequest, SenderAddr: from.String()}
	err := SendMessage(conn, message)
	if err != nil {
		return
	}

	err = conn.Close()
	peerMap.AddPeer(Peer{Addr: to})
}

func SendRemoveMeRequest(from net.Addr, to net.Addr) {
	fmt.Println("disconnecting  from:", to)
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
