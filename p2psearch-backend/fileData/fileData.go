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
	Mutex sync.RWMutex
	Data  map[string]FileData
}

func CreateFileDataStore() FileDataStore {
	return FileDataStore{
		Mutex: sync.RWMutex{},
		Data:  make(map[string]FileData),
	}
}

func (store *FileDataStore) HasFileStored(path string) bool {
	store.Mutex.RLock()
	_, hasFile := store.Data[path]
	store.Mutex.RUnlock()
	return hasFile
}

func (store *FileDataStore) RetrieveFileData(path string) FileData {
	store.Mutex.RLock()
	fileData, _ := store.Data[path]
	store.Mutex.RUnlock()
	return fileData
}

func (store *FileDataStore) StoreFileData(fileData FileData) {
	store.Mutex.Lock()
	store.Data[fileData.Url] = fileData
	store.Mutex.Unlock()
}
