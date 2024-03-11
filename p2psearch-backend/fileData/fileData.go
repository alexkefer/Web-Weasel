package fileData

import (
	"sync"
	"time"
)

type FileData struct {
	Url          string
	FileLoc      string
	FileType     string
	DownloadTime time.Time
	AccessTime   time.Time
}

func CreateFileData(url string, fileLoc string, fileType string) FileData {
	return FileData{
		Url:          url,
		FileLoc:      fileLoc,
		FileType:     fileType,
		DownloadTime: time.Now(),
		AccessTime:   time.Now(),
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
	store.data[fileData.Url] = fileData
	store.mutex.Unlock()
}
