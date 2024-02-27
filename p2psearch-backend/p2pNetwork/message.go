package p2pNetwork

import (
	"encoding/gob"
	"github.com/alexkefer/p2psearch-backend/log"
	"io"
	"net"
)

type Message struct {
	Code             int
	SenderAddr       string
	Peers            []string
	BroadcastMessage string
}

func ReceiveMessage(r io.Reader) (Message, error) {
	decoder := gob.NewDecoder(r)

	var message Message

	err := decoder.Decode(&message)
	if err != nil {
		log.Error("failed receiving p2p message: %s", err)
	}
	return message, err
}

func SendMessage(w io.Writer, message Message) error {
	encoder := gob.NewEncoder(w)
	err := encoder.Encode(message)

	if err != nil {
		log.Error("failed sending p2p message: %s", err)
	}

	return err
}

func NetAddrMapToStringMap(netAddrMap map[net.Addr]int) map[string]int {
	stringMap := make(map[string]int)

	for addr, _ := range netAddrMap {
		stringMap[addr.String()] = 0
	}

	return stringMap
}

func StringMapToNetAddrMap(stringMap map[string]int) map[net.Addr]int {
	netAddrMap := make(map[net.Addr]int)

	for addr, _ := range stringMap {
		addr, addrParseErr := net.ResolveTCPAddr("tcp", addr)

		if addrParseErr != nil {
			panic(addrParseErr)
		}

		netAddrMap[addr] = 0
	}

	return netAddrMap
}
