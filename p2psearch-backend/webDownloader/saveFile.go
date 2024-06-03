/* Keagan Edwards & Alex Kefer - SaveFile function */

package webDownloader

import (
	"github.com/alexkefer/p2psearch-backend/fileData"
	"github.com/alexkefer/p2psearch-backend/log"
	"github.com/alexkefer/p2psearch-backend/utils"
	"os"
	"path/filepath"
)

/*
SaveFile saves the content of the file inside of a filename with its filetype at the fileDataStore location.
The function creates its required folder and each page required to load the webpage.
*/
func SaveFile(content []byte, filename string, fileType string, fileDataStore *fileData.FileDataStore) {
	saveLocation, err2 := utils.GetCachePath()
	fullSaveLocation := filepath.Join(saveLocation, filename)
	log.Debug("saving asset: %s, %s, %s", filename, fullSaveLocation, fileType)

	if err2 != nil {
		log.Error("failed to save page: %s", err2)
		return
	}

	err := os.MkdirAll(saveLocation, os.ModePerm)

	if err != nil {
		log.Error("error saving asset: %s", err)
		return
	}
	// takes in the context of the page and saves it to the save location
	file, err := os.OpenFile(fullSaveLocation, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Error("error opening file: %s", err)
		return
	}
	defer file.Close()

	_, err3 := file.Write(content)
	if err3 != nil {
		log.Error("error writing to file: %s", err3)
		return
	} else {
		log.Info("successfully saved file: %s", filename)
	}

	metadata := fileData.CreateFileData(filename, fullSaveLocation, fileType, true)
	fileDataStore.StoreFileData(metadata)
}
