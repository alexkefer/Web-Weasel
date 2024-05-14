// Package fileData is responsible for tracking cached resources in memory during runtime.
//
// Every cached resource should hava a corresponding FileData struct, stored in a FileDataStore struct.
package fileData

import (
	"encoding/json"
	"github.com/alexkefer/p2psearch-backend/log"
	"github.com/alexkefer/p2psearch-backend/utils"
	"os"
	"sync"
	"time"
)

// FileData enumerates metadata for a single cached resource.
type FileData struct {
	// Url describes the address a resource was received from.
	Url string
	// FileLoc is the filepath pointing to cached resource.
	FileLoc string
	// FileType is the resources MIME type. This is used by the HTTP server when serving the cached resource.
	FileType string
	// DownloadTime is the time the resource was cached.
	DownloadTime time.Time
	// AccessTime is the last time this file was accessed by an HTTP client or another peer.
	AccessTime time.Time
}

// CreateFileData creates a new FileData struct from the provided arguments. It initializes the FileData.DownloadTime
// and FileData.AccessTime fields to the current time.
func CreateFileData(url string, fileLoc string, fileType string) FileData {
	return FileData{
		Url:          url,
		FileLoc:      fileLoc,
		FileType:     fileType,
		DownloadTime: time.Now(),
		AccessTime:   time.Now(),
	}
}

// FileDataStore holds a map which holds FileData structs for every cached resource.
type FileDataStore struct {
	Mutex sync.RWMutex
	Data  map[string]FileData
}

// CreateFileDataStore initializes a new FileDataStore struct.
func CreateFileDataStore() FileDataStore {
	return FileDataStore{
		Mutex: sync.RWMutex{},
		Data:  make(map[string]FileData),
	}
}

// HasFileStored checks if the corresponding FileDataStore contains a FileData struct with a FileData.Url field
// corresponding to path.
func (store *FileDataStore) HasFileStored(path string) bool {
	store.Mutex.RLock()
	_, hasFile := store.Data[path]
	store.Mutex.RUnlock()
	return hasFile
}

// RetrieveFileData return a copy of the FileData struct with a FileData.Url field corresponding to path, stored in the
// FileDataStore struct.
func (store *FileDataStore) RetrieveFileData(path string) FileData {
	store.Mutex.RLock()
	fileData, _ := store.Data[path]
	store.Mutex.RUnlock()
	return fileData
}

// StoreFileData stores a FileData struct in the FileDataStore, using the FileData.Url field as the key.
func (store *FileDataStore) StoreFileData(fileData FileData) {
	store.Mutex.Lock()
	store.Data[fileData.Url] = fileData
	store.Mutex.Unlock()
}

// SaveFileDataStore saves the FileDataStore struct to a metadata file in the cache directory using JSON format.
func (store *FileDataStore) SaveFileDataStore() {
	log.Debug("saving resource metadata")
	store.Mutex.RLock()

	cachePath, err := utils.GetCachePath()

	if err == nil {
		file, fileErr := os.OpenFile(
			cachePath+string(os.PathSeparator)+"metadata.json",
			os.O_RDWR|os.O_CREATE,
			0644)

		if fileErr == nil {
			encoder := json.NewEncoder(file)
			encodeErr := encoder.Encode(store.Data)
			if encodeErr != nil {
				log.Error("problem encoding Data store to metadata file: %s", encodeErr)
			}
			closeErr := file.Close()
			if closeErr != nil {
				log.Warn("problem closing metadata file: %d", closeErr)
			}
		} else {
			log.Error("problem opening metadata file for saving: %s", fileErr)
		}
	}

	store.Mutex.RUnlock()
}

// LoadFileData loads resource metadata from the JSON file in the cache directory if it exists into the FileDataStore.
// This function overwrites any Data that is already in the FileDataStore.
func (store *FileDataStore) LoadFileData() {
	store.Mutex.Lock()

	log.Debug("loading resource metadata from file")

	cachePath, err := utils.GetCachePath()
	if err == nil {
		file, fileErr := os.OpenFile(
			cachePath+string(os.PathSeparator)+"metadata.json",
			os.O_RDONLY,
			0644)

		if fileErr == nil {
			decoder := json.NewDecoder(file)
			decodeErr := decoder.Decode(&store.Data)
			if decodeErr != nil {
				log.Warn("problem decoding metadata file: %s", decodeErr)
			}
		} else {
			log.Warn("problem opening metadata file: %s", fileErr)
		}
	}

	store.Mutex.Unlock()
}
