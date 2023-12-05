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
func ParseCommands(addressChan chan<- net.Addr, myAddr net.Addr, connectedParties map[net.Addr]int) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Read a command from the user
		fmt.Print(">> ")
		if !scanner.Scan() {
			break
		}
		command := scanner.Text()

		switch command {
		case "list":
			ListConnections(connectedParties)
		case "broadcast":
			fmt.Print("[enter message]: ")
			if !scanner.Scan() {
				break
			}
			message := scanner.Text()

			for toAddr, _ := range connectedParties {
				SendBroadcastMessage(toAddr, myAddr, message)
			}

		case "msg":
			fmt.Print("[enter target address]: ")
			if !scanner.Scan() {
				break
			}
			addressStr := scanner.Text()

			toAddr, addrParseErr := net.ResolveTCPAddr("tcp", addressStr)
			if addrParseErr != nil {
				fmt.Println("address parse error:", addrParseErr)
				continue
			}

			fmt.Print("[enter message]: ")
			if !scanner.Scan() {
				break
			}
			message := scanner.Text()

			SendBroadcastMessage(toAddr, myAddr, message)

		case "exit":
			fmt.Println("WIP")
		case "help":
			Help()
		default:
			fmt.Println("invalid command")
		}
	}
}

func ListConnections(addresses map[net.Addr]int) {
	fmt.Printf("Connected Parties: ")
	for addr, _ := range addresses {
		fmt.Printf("%s, ", addr)
	}
	fmt.Println()
}

// Help -- Displays help
func Help() {
	fmt.Println("Commands:")
	fmt.Println("  - list")
	fmt.Println("  - broadcast <message> -- broadcast a message to all connected parties")
	fmt.Println("  - msg <ip address>:<port> <message> -- send a message to a specific address")
	fmt.Println("  - exit -- exit the program")
	fmt.Println("  - help -- display help")
}
