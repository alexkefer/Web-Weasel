package main

import (
	"fmt"
	"net"
)

func StoreAddresses(addressChan <-chan net.Addr) {
	// Here we are creating a map, go's equivalent of dictionaries in Python or HashMap in Java
	addresses := make(map[net.Addr]int)

	for {
		address := <-addressChan
		addresses[address] = 0

		fmt.Printf("address added to store: %s, addr store size: %d\n", address, len(addresses))
	}
}
