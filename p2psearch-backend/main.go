package main

import (
	"fmt"
	"github.com/alexkefer/p2psearch-backend/fileData"
	"github.com/alexkefer/p2psearch-backend/httpServer"
	"github.com/alexkefer/p2psearch-backend/p2pServer"
	"github.com/alexkefer/p2psearch-backend/utils"
	"net"
	"os"
	"os/signal"
)

// This is the main function of the program
func main() {

	if len(os.Args) > 2 {
		fmt.Println("error: either 0 or 1 arguments expected")
		return
	}
	port, portErr := utils.FindOpenPort(9000, 9100)
	if portErr != nil {
		fmt.Println("error finding open port:", portErr)
		return
	}
	address := utils.GetLocalIPAddress() + port
	myAddr, myAddrParseErr := net.ResolveTCPAddr("tcp", address)

	if myAddrParseErr != nil {
		fmt.Println("error parsing my ip arg:", myAddrParseErr)
		return
	}

	peerMap := p2pServer.PeerMap{Peers: make(map[string]p2pServer.Peer)}
	myPeer := p2pServer.Peer{Addr: myAddr}
	peerMap.AddPeer(myPeer)

	go p2pServer.RequestHandler(myAddr, &peerMap)
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
			p2pServer.SendAddMeRequest(myAddr, seedAddr, &peerMap)
			p2pServer.SendMoreAddMeRequests(myAddr, seedAddr, &peerMap)
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

	fileDataStore := fileData.CreateFileDataStore()

	go httpServer.StartServer(&peerMap, &fileDataStore, exitChannel)

	for {
		if <-exitChannel {
			break
		}
	}

	peerMap.Mutex.RLock()
	for _, peer := range peerMap.Peers {
		if peer.Addr != myAddr {
			p2pServer.SendRemoveMeRequest(myAddr, peer.Addr)
		}
	}
	peerMap.Mutex.RUnlock()
}
