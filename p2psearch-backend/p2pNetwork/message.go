package p2pNetwork

import (
	"encoding/gob"
	"github.com/alexkefer/p2psearch-backend/log"
	"io"
)

type Message struct {
	Code       int
	SenderAddr string
	DataType   string
	Peers      []string
	Data       []byte
}

const (
	AddMeRequest = iota
	SharePeersRequest
	SharePeersResponse
	BroadcastMessage
	RemoveMeRequest
	FileRequest
	HasFileResponse
	NoFileResponse
)

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
