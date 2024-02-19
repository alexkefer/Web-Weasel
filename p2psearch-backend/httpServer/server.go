package httpServer

import (
	"github.com/alexkefer/p2psearch-backend/fileData"
	"github.com/alexkefer/p2psearch-backend/log"
	"github.com/alexkefer/p2psearch-backend/p2pNetwork"
	"github.com/alexkefer/p2psearch-backend/utils"
	"net"
	"net/http"
)

func StartServer(peerMap *p2pNetwork.PeerMap, fileDataStore *fileData.FileDataStore, shutdownChan chan<- bool, myAddr net.Addr) {
	http.HandleFunc("/", defaultHandler)

	http.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		shutdownHandler(w, r, shutdownChan)
	})

	http.HandleFunc("/peers", func(w http.ResponseWriter, r *http.Request) {
		peersHandler(w, r, peerMap)
	})

	http.HandleFunc("/cache", func(w http.ResponseWriter, r *http.Request) {
		cacheFileHandler(w, r, fileDataStore)
	})

	http.HandleFunc("/retrieve", func(w http.ResponseWriter, r *http.Request) {
		retrieveFileHandler(w, r, fileDataStore)
	})

	http.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		connectHandler(w, r, myAddr, peerMap)
	})

	http.HandleFunc("/disconnect", func(w http.ResponseWriter, r *http.Request) {
		disconnectHandler(w, myAddr, peerMap)
	})

	port, _ := utils.FindOpenPort(8080, 8180)

	log.Info("opening http server on port %s", port)

	http.ListenAndServe(port, nil)
}
