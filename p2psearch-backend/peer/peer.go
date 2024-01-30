package peer

import (
	"net"
	"sync"
)

type PeerMap struct {
	Mutex sync.RWMutex
	Peers map[string]Peer
}

func (peerMap *PeerMap) HasPeer(addr string) bool {
	peerMap.Mutex.RLock()
	_, hasPeer := peerMap.Peers[addr]
	peerMap.Mutex.RUnlock()
	return hasPeer
}

func (peerMap *PeerMap) AddPeer(peer Peer) {
	peerMap.Mutex.Lock()
	key := peer.Addr.String()
	peerMap.Peers[key] = peer
	peerMap.Mutex.Unlock()

}

func (peerMap *PeerMap) RemovePeer(addr string) {
	peerMap.Mutex.Lock()
	delete(peerMap.Peers, addr)
	peerMap.Mutex.Unlock()

}

type Peer struct {
	Addr net.Addr
}
