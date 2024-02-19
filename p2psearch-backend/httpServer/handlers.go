package httpServer

import (
	"fmt"
	"github.com/alexkefer/p2psearch-backend/fileData"
	"github.com/alexkefer/p2psearch-backend/log"
	"github.com/alexkefer/p2psearch-backend/p2pNetwork"
	"github.com/alexkefer/p2psearch-backend/utils"
	"github.com/alexkefer/p2psearch-backend/webDownloader"
	"html"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func defaultHandler(w http.ResponseWriter, r *http.Request, fileDataStore *fileData.FileDataStore) {
	path := utils.UrlCleaner(r.URL.Path)
	path = strings.ReplaceAll(path, "/", "_")

	log.Info("path: %q", path)

	retrieveFile(w, path, fileDataStore)
}

func shutdownHandler(w http.ResponseWriter, r *http.Request, shutdownChan chan<- bool) {
	shutdownChan <- true
	fmt.Fprintf(w, "shutting down server...")
}

func peersHandler(w http.ResponseWriter, r *http.Request, peerMap *p2pNetwork.PeerMap) {
	peerMap.Mutex.RLock()
	for key, _ := range peerMap.Peers {
		fmt.Fprintf(w, "%s\n", html.EscapeString(key))
	}
	peerMap.Mutex.RUnlock()
}

func cacheFileHandler(w http.ResponseWriter, r *http.Request, fileDataStore *fileData.FileDataStore) {
	path, err := getPathParam(r.URL)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Warn("error caching file: %s", err)
		fmt.Fprintf(w, "error caching file: %s", err)
	} else {
		err = webDownloader.BuildDownloadedWebpage(path, fileDataStore)
		if err == nil {
			fmt.Fprintf(w, "server cached resource found at %q", path)
		} else {
			fmt.Fprintf(w, "server failed to cache resource: %s", err)
		}
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
		metadata := fileData.RetrieveFileData(path)

		file, openErr := os.Open(metadata.FileLoc)

		if openErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "server couldn't open file for: %q", path)
			log.Error("couldn't open file %q: %s", metadata.FileLoc, openErr)
			return
		}

		_, fileErr := file.WriteTo(w)

		if fileErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error("couldn't read file %q: %s", metadata.FileLoc, fileErr)
			return
		}

		w.Header().Add("Content-Type", metadata.FileType)
	} else {
		// TODO: Ask other hosts on the network if they have it.
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "server has no resource associated with path: %q", path)
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

	connectErr := p2pNetwork.Connect(myAddr, targetAddr, peerMap)

	if connectErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "failed to connecting to peer: %s", connectErr)
		return
	}

	fmt.Fprintf(w, "sent add me request to: %s", targetAddr.String())
	log.Info("sent add me request to: %s", targetAddr.String())
}

func disconnectHandler(w http.ResponseWriter, myAddr net.Addr, peerMap *p2pNetwork.PeerMap) {
	log.Info("received http disconnect request")
	p2pNetwork.Disconnect(myAddr, peerMap)
	fmt.Fprintf(w, "disconnecting from p2p network")
}

func filesHandler(w http.ResponseWriter, fileDataStore *fileData.FileDataStore) {
	fileDataStore.Mutex.RLock()
	for key, _ := range fileDataStore.Data {
		fmt.Fprintf(w, "%s\n", key)
	}
	fileDataStore.Mutex.RUnlock()
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

func retrieveFile(w http.ResponseWriter, cleanUrl string, fileData *fileData.FileDataStore) {

	if fileData.HasFileStored(cleanUrl) {
		metadata := fileData.RetrieveFileData(cleanUrl)

		file, openErr := os.Open(metadata.FileLoc)

		if openErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "server couldn't open file for: %q", cleanUrl)
			log.Error("couldn't open file %q: %s", metadata.FileLoc, openErr)
			return
		}

		_, fileErr := file.WriteTo(w)

		if fileErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error("couldn't read file %q: %s", metadata.FileLoc, fileErr)
			return
		}

		w.Header().Add("Content-Type", metadata.FileType)
	} else {
		// TODO: Ask other hosts on the network if they have it.
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "server has no resource associated with path: %q", cleanUrl)
	}
}
