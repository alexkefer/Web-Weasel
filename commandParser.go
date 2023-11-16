// 11-15-2023 - Alex Kefer - commandParser.go
// Parses the commands from user input after the program has started
// Commands are:
//  1. msgAll <message> - sends a message to all connected parties
//  2. msg <address> <message> - sends a message to the specified address
//  3. list - lists all connected parties
//  4. exit - exits the program
//  5. help - lists all commands
package main

import (
	"bufio"
	"fmt"
	"os"
)

func PromptUser() string {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		switch text {
		case "msgAll":
		case "msg":
		case "list":
		case "exit":
		}
		fmt.Println("text:", text)
	}
}
