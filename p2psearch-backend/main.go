// Main package for the peer-to-peer web cache backend. This package contains the entry point for the backend.
package main

import (
	"github.com/alexkefer/p2psearch-backend/fileData"
	"github.com/alexkefer/p2psearch-backend/httpServer"
	"github.com/alexkefer/p2psearch-backend/log"
	"github.com/alexkefer/p2psearch-backend/p2pNetwork"
	"github.com/alexkefer/p2psearch-backend/utils"
	"net"
	"os"
	"os/signal"
)

// main starts the peer-to-peer backend and http web server.
func main() {

	if len(os.Args) > 2 {
		log.Error("either 0 or 1 arguments expected")
		return
	}
	port, portErr := utils.FindOpenPort(9000, 9100)
	if portErr != nil {
		log.Error("error finding open port: %s", portErr)
		return
	}
	address := utils.GetLocalIPAddress() + port
	myAddr, myAddrParseErr := net.ResolveTCPAddr("tcp", address)

	if myAddrParseErr != nil {
		log.Error("couldn't parsing my ip arg:", myAddrParseErr)
		return
	}

	peerMap := p2pNetwork.PeerMap{Peers: make(map[string]p2pNetwork.Peer)}
	myPeer := p2pNetwork.Peer{Addr: myAddr}
	peerMap.AddPeer(myPeer)

	fileDataStore := fileData.CreateFileDataStore()

	go p2pNetwork.StartServer(myAddr, &peerMap, &fileDataStore)
	log.Info("my address: %s", myAddr)

	// If an address is given, try to join its network
	if len(os.Args) > 1 {
		seedAddrArg := os.Args[1]

		seedAddr, addrParseErr := net.ResolveTCPAddr("tcp", seedAddrArg)

		if addrParseErr != nil {
			log.Error("seedAddr parse error:", addrParseErr)
			return
		} else {
			connectErr := p2pNetwork.Connect(myAddr, seedAddr, &peerMap)

			if connectErr != nil {
				return
			}
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

	go httpServer.StartServer(&peerMap, &fileDataStore, exitChannel, myAddr)

	for {
		if <-exitChannel {
			break
		}
	}

	p2pNetwork.Disconnect(myAddr, &peerMap)
}
