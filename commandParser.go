/*
	Description: This file contains the command parser for the p2p network
	Communicates with the RequestHandler on the other end of the connection to send commands

Commands:
  - list
  - broadcast <message> -- broadcast a message to all connected parties
  - msg <ip address>:<port> <message> -- send a message to a specific address
  - exit -- exit the program
  - help -- display help
    Future Potential Commands:
  - connect <ip address>:<port> -- potential future feature, maybe be able to "merge" networks
  - disconnect -- potential future feature, maybe be to leave a network to join another
*/
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	ListCommand      = 0
	BroadcastCommand = 1
	MessageCommand   = 2
	ExitCommand      = 3
	HelpCommand      = 4
)

// ParseCommands -- Goroutine that parses commands from the user, and sends them to the RequestHandler, automatically resets the prompt after handling a request
func ParseCommands(myAddr net.Addr, peerMap *PeerMap) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Read a command from the user
		fmt.Print(">> ")
		if !scanner.Scan() {
			return
		}
		command := scanner.Text()

		switch command {
		case "list":
			ListConnections(peerMap)
		case "broadcast":
			fmt.Print("[enter message]: ")
			if !scanner.Scan() {
				return
			}
			message := scanner.Text()

			peerMap.mutex.RLock()
			for _, peer := range peerMap.peers {
				SendBroadcastMessage(peer.addr, myAddr, message)
			}
			peerMap.mutex.RUnlock()

		case "msg":
			fmt.Print("[enter target address]: ")
			if !scanner.Scan() {
				return
			}
			addressStr := scanner.Text()

			toAddr, addrParseErr := net.ResolveTCPAddr("tcp", addressStr)
			if addrParseErr != nil {
				fmt.Println("address parse error:", addrParseErr)
				continue
			}

			fmt.Print("[enter message]: ")
			if !scanner.Scan() {
				return
			}
			message := scanner.Text()

			SendBroadcastMessage(toAddr, myAddr, message)

		case "exit":
			return
		case "help":
			Help()
		default:
			fmt.Println("invalid command")
		}
	}
}

func ListConnections(peerMap *PeerMap) {
	peerMap.mutex.RLock()
	fmt.Printf("Connected Parties: ")
	for addr, _ := range peerMap.peers {
		fmt.Printf("%s, ", addr)
	}
	fmt.Println()
	peerMap.mutex.RUnlock()
}

// Help -- Displays help
func Help() {
	fmt.Println("Commands:")
	fmt.Println("  - list")
	fmt.Println("  - broadcast -- broadcast a message to all connected parties")
	fmt.Println("  - msg -- send a message to a specific address")
	fmt.Println("  - exit -- exit the program")
	fmt.Println("  - help -- display help")
}
