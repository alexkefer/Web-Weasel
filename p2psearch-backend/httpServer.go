package main

import (
	"fmt"
	"github.com/alexkefer/p2psearch-backend/peer"
	"github.com/alexkefer/p2psearch-backend/utils"
	"html"
	"net/http"
	"net/url"
)

func RunHttpServer(peerMap *peer.PeerMap, shutdownChan chan<- bool) {
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
		retrieveFileHandler(w, r)
	})

	port, _ := utils.FindOpenPort(8080, 8180)

	fmt.Printf("opening http server on port %s\n", port)

	http.ListenAndServe(port, nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.String()))
}

func shutdownHandler(w http.ResponseWriter, r *http.Request, shutdownChan chan<- bool) {
	shutdownChan <- true
	fmt.Fprintf(w, "Shutting down server...")
}

func peersHandler(w http.ResponseWriter, r *http.Request, peerMap *peer.PeerMap) {
	peerMap.Mutex.RLock()
	for key, _ := range peerMap.Peers {
		fmt.Fprintf(w, "%s\n", html.EscapeString(key))
	}
	peerMap.Mutex.RUnlock()
}

func storeFileHandler(w http.ResponseWriter, r *http.Request) {
	path, err := getPathParam(r.URL)

	if err != nil {
		fmt.Fprintf(w, "Error storing file: %s", err)
	} else {
		// TODO
		fmt.Fprintf(w, "Server is storing file found at %q", path)
	}
}

func retrieveFileHandler(w http.ResponseWriter, r *http.Request) {
	path, err := getPathParam(r.URL)

	if err != nil {
		fmt.Fprintf(w, "Error retrieving file: %s", err)
	} else {
		// TODO
		fmt.Fprintf(w, "Server is retrieving file found at %q", path)
	}
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
