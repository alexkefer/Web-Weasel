package httpServer

import (
	"fmt"
	"github.com/alexkefer/p2psearch-backend/fileData"
	"github.com/alexkefer/p2psearch-backend/log"
	"github.com/alexkefer/p2psearch-backend/p2pNetwork"
	"html"
	"net"
	"net/http"
	"net/url"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.String()))
}

func shutdownHandler(w http.ResponseWriter, r *http.Request, shutdownChan chan<- bool) {
	shutdownChan <- true
	fmt.Fprintf(w, "Shutting down server...")
}

func peersHandler(w http.ResponseWriter, r *http.Request, peerMap *p2pNetwork.PeerMap) {
	peerMap.Mutex.RLock()
	for key, _ := range peerMap.Peers {
		fmt.Fprintf(w, "%s\n", html.EscapeString(key))
	}
	peerMap.Mutex.RUnlock()
}

func storeFileHandler(w http.ResponseWriter, r *http.Request) {
	path, err := getPathParam(r.URL)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error storing file: %s", err)
	} else {
		// TODO
		fmt.Fprintf(w, "Server is storing file found at %q", path)
	}
}

func retrieveFileHandler(w http.ResponseWriter, r *http.Request, fileData *fileData.FileDataStore) {
	path, err := getPathParam(r.URL)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error retrieving file: %s", err)
		return
	}

	if fileData.HasFileStored(path) {
		// TODO
	} else {
		// TODO: Ask other hosts on the network if they have it.
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Server has no resource associated with path: %q", path)
	}
}

func connectHandler(w http.ResponseWriter, r *http.Request, myAddr net.Addr, peerMap *p2pNetwork.PeerMap) {
	path, err := getPathParam(r.URL)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "failed to connecting to peer: %s", err)
		log.Warn("failed to connect to peer: %s", err)
		return
	}

	targetAddr, addrParseErr := net.ResolveTCPAddr("tcp", path)

	if addrParseErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "failed to connecting to peer: %s", addrParseErr)
		log.Warn("failed to connect to peer: %s", addrParseErr)
		return
	}

	addMeError := p2pNetwork.SendAddMeRequest(myAddr, targetAddr, peerMap)

	if addMeError != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "failed to connecting to peer: %s", addMeError)
		log.Error("failed to connect to peer: %s", addMeError)
		return
	}

	fmt.Fprintf(w, "sent add me request to: %s", targetAddr.String())
	log.Info("sent add me request to: %s", targetAddr.String())
}

func getPathParam(fromUrl *url.URL) (string, error) {
	params, err := url.ParseQuery(fromUrl.RawQuery)

	if err != nil {
		return "", err
	}

	if params.Has("path") {
		return params.Get("path"), nil
	}

	return "", fmt.Errorf("request has no path param")
}
