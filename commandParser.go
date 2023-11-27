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
	"fmt"
	"net"
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
	for {
		// Read a command from the user
		var command string
		fmt.Print(">> ")
		fmt.Scanln(&command)

		switch command {
		case "list":
			ListConnections(connectedParties)
		case "broadcast":
			var message string
			fmt.Scanln(&message)
			fmt.Println("WIP")
		case "msg":
			var address string
			fmt.Scanln(&address)
			var message string
			fmt.Scanln(&message)
			fmt.Println("WIP")
		case "exit":
			fmt.Println("WIP")
		case "help":
			Help()
		default:
			fmt.Println("Invalid command")
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
