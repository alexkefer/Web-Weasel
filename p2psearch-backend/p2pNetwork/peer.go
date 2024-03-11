package p2pNetwork

import (
	"github.com/alexkefer/p2psearch-backend/log"
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
	log.Debug("adding peer: %s", peer.Addr.String())

	peerMap.Mutex.Lock()
	key := peer.Addr.String()
	peerMap.Peers[key] = peer
	peerMap.Mutex.Unlock()

}

func (peerMap *PeerMap) RemovePeer(addr string) {

	if !peerMap.HasPeer(addr) {
		log.Warn("cannot remove peer that does not exist: %s", addr)
		return
	}

	log.Debug("removing peer: %s", addr)

	peerMap.Mutex.Lock()
	delete(peerMap.Peers, addr)
	peerMap.Mutex.Unlock()

}

type Peer struct {
	Addr net.Addr
}
