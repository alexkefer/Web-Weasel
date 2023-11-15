package main

import (
	"fmt"
	"net"
)

func StoreAddresses(addresses map[net.Addr]int, addressChan <-chan net.Addr) {
	// Here we are creating a map, go's equivalent of dictionaries in Python or HashMap in Java
	for {
		address := <-addressChan
		addresses[address] = 0
		fmt.Printf("address added to store: %s, addr store size: %d\n", address, len(addresses))
		fmt.Printf("Entire Network: %v\n", addresses)
	}
}
