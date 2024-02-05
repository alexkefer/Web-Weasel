package fileData

import (
	"sync"
	"time"
)

type FileData struct {
	path         string
	fileType     string
	downloadTime time.Time
	accessTime   time.Time
}

func CreateFileData(path string, fileType string) FileData {
	return FileData{
		path:         path,
		fileType:     fileType,
		downloadTime: time.Now(),
		accessTime:   time.Now(),
	}
}

type FileDataStore struct {
	mutex sync.RWMutex
	data  map[string]FileData
}

func CreateFileDataStore() FileDataStore {
	return FileDataStore{
		mutex: sync.RWMutex{},
		data:  make(map[string]FileData),
	}
}

func (store *FileDataStore) HasFileStored(path string) bool {
	store.mutex.RLock()
	_, hasFile := store.data[path]
	store.mutex.RUnlock()
	return hasFile
}

func (store *FileDataStore) RetrieveFileData(path string) FileData {
	store.mutex.RLock()
	fileData, _ := store.data[path]
	store.mutex.RUnlock()
	return fileData
}

func (store *FileDataStore) StoreFileData(fileData FileData) {
	store.mutex.Lock()
	store.data[fileData.path] = fileData
	store.mutex.Unlock()
}
