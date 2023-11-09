package main

import (
	"encoding/json"
	"net"
)

type Message struct {
	Code       int
	SenderAddr string
}

func ReceiveMessage(conn net.Conn) (Message, error) {
	decoder := json.NewDecoder(conn)

	var message Message

	err := decoder.Decode(&message)

	return message, err
}

func SendMessage(conn net.Conn, message Message) error {
	encoder := json.NewEncoder(conn)
	return encoder.Encode(message)
}
