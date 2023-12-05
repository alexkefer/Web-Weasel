package main

import (
	"net"
	"sync"
)

type PeerMap struct {
	mutex sync.RWMutex
	peers map[string]Peer
}

func (peerMap *PeerMap) HasPeer(addr string) bool {
	peerMap.mutex.RLock()
	_, hasPeer := peerMap.peers[addr]
	peerMap.mutex.RUnlock()
	return hasPeer
}

func (peerMap *PeerMap) AddPeer(peer Peer) {
	peerMap.mutex.Lock()
	key := peer.addr.String()
	peerMap.peers[key] = peer
	peerMap.mutex.Unlock()

}

type Peer struct {
	addr net.Addr
}
