package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type Message struct {
	Code             int
	SenderAddr       string
	Peers            []string
	BroadcastMessage string
}

func ReceiveMessage(conn net.Conn) (Message, error) {
	decoder := json.NewDecoder(conn)

	var message Message

	err := decoder.Decode(&message)
	if err != nil {
		fmt.Println("error receiving message:", err)
	}
	return message, err
}

func SendMessage(conn net.Conn, message Message) error {
	encoder := json.NewEncoder(conn)
	err := encoder.Encode(message)

	if err != nil {
		fmt.Println("error sending message:", err)
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
