package main

import (
	"fmt"
	"html"
	"net/http"
)

func RunHttpServer(peerMap *PeerMap, shutdownChan chan<- bool) {
	http.HandleFunc("/", helloHandler)

	http.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		shutdownHandler(w, r, shutdownChan)
	})

	http.HandleFunc("/peers", func(w http.ResponseWriter, r *http.Request) {
		peersHandler(w, r, peerMap)
	})

	port, _ := findOpenPort(8080, 8180)

	fmt.Printf("opening http server on port %s\n", port)
	
	http.ListenAndServe(port, nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func shutdownHandler(w http.ResponseWriter, r *http.Request, shutdownChan chan<- bool) {
	shutdownChan <- true
	fmt.Fprintf(w, "Shutting down server...")
}

func peersHandler(w http.ResponseWriter, r *http.Request, peerMap *PeerMap) {
	peerMap.mutex.RLock()
	for key, _ := range peerMap.peers {
		fmt.Fprintf(w, "%s\n", html.EscapeString(key))
	}
	peerMap.mutex.RUnlock()
}
