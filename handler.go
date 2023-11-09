package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func RequestHandler(myAddr net.Addr, addressChan chan<- net.Addr) {
	listener, listenErr := net.Listen("tcp", myAddr.String())

	if listenErr != nil {
		fmt.Println("tcp listen error:", listenErr)
		return
	}

	// This is an infinite loop in go
	for {
		conn, connErr := listener.Accept()

		if connErr != nil {
			fmt.Println("tcp connection error:", connErr)
		} else {
			// Here we create a separate goroutine (thread) to handle this connection
			go HandleConnection(conn, addressChan)
		}
	}
}

func HandleConnection(conn net.Conn, addressChan chan<- net.Addr) {
	fmt.Println("connection received from:", conn.RemoteAddr())
	reader := bufio.NewReader(conn)
	codeStr, readErr := reader.ReadString('\n')

	if readErr != nil {
		fmt.Println("request parse error:", readErr)
		CloseConnection(conn)
		return
	}

	code, codeParseErr := strconv.Atoi(strings.TrimSpace(codeStr))

	if codeParseErr != nil {
		fmt.Println("request code parse error:", codeParseErr)
		CloseConnection(conn)
		return
	}

	fmt.Printf("code: %d\n", code)

	switch code {
	case JoinRequest:

		addrStr, readErr := reader.ReadString('\n')

		if readErr != nil {
			fmt.Println("request parse error:", readErr)
			CloseConnection(conn)
			return
		}

		addrStr = strings.TrimSpace(addrStr)
		fmt.Println("addr str:", addrStr)

		addr, addrParseErr := net.ResolveTCPAddr("tcp", addrStr)

		if addrParseErr != nil {
			fmt.Println("addr parse error:", addrParseErr)
			return
		}

		addressChan <- addr
	default:
		fmt.Printf("invalid code %d, closing connection.\n", code)

	}

	CloseConnection(conn)
}

func CloseConnection(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		fmt.Println("connection close error:", err)
	} else {
		fmt.Println("closed connection:", conn.RemoteAddr())
	}
}
