package httpServer

import (
	"github.com/alexkefer/p2psearch-backend/fileData"
	"github.com/alexkefer/p2psearch-backend/log"
	"github.com/alexkefer/p2psearch-backend/p2pServer"
	"github.com/alexkefer/p2psearch-backend/utils"
	"net/http"
)

func StartServer(peerMap *p2pServer.PeerMap, fileData *fileData.FileDataStore, shutdownChan chan<- bool) {
	http.HandleFunc("/", helloHandler)

	http.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		shutdownHandler(w, r, shutdownChan)
	})

	http.HandleFunc("/peers", func(w http.ResponseWriter, r *http.Request) {
		peersHandler(w, r, peerMap)
	})

	http.HandleFunc("/store", func(w http.ResponseWriter, r *http.Request) {
		storeFileHandler(w, r)
	})

	http.HandleFunc("/retrieve", func(w http.ResponseWriter, r *http.Request) {
		retrieveFileHandler(w, r, fileData)
	})

	port, _ := utils.FindOpenPort(8080, 8180)

	log.Info("opening http server on port %s", port)

	http.ListenAndServe(port, nil)
}
