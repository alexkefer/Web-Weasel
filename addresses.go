package main

import (
	"net"
)

func StoreAddresses(addresses map[net.Addr]int, addressChan <-chan net.Addr, finishedStore chan<- bool) {
	// Here we are creating a map, go's equivalent of dictionaries in Python or HashMap in Java
	for {
		address := <-addressChan
		addresses[address] = 0
		finishedStore <- true
	}
}
