package utils

import (
	"fmt"
	"net"
	"time"
)

func MakeTcpConnection(to net.Addr) (net.Conn, error) {
	duration, parseErr := time.ParseDuration("5s")

	if parseErr != nil {
		fmt.Println("duration parse error:", parseErr)
		return nil, parseErr
	}

	conn, connErr := net.DialTimeout("tcp", to.String(), duration)

	if connErr != nil {
		fmt.Println("error connecting to address:", connErr)
		return nil, connErr
	}

	return conn, nil
}
