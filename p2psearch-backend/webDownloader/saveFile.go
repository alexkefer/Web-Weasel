/* Function and helpers to save pages to file at specified location*/

package webDownloader

import (
	"github.com/alexkefer/p2psearch-backend/fileData"
	"github.com/alexkefer/p2psearch-backend/log"
	"github.com/alexkefer/p2psearch-backend/utils"
	"os"
	"path/filepath"
	"strings"
)

func savePage(content []byte, url string, fileType string, fileDataStore *fileData.FileDataStore) {
	urlClean := utils.UrlCleaner(url)
	//fileExt := fileTypes.GetFileExtension(fileType)

	urlClean = strings.ReplaceAll(urlClean, "/", "_")

	saveLocation, err2 := utils.GetCachePath()
	fullSaveLocation := filepath.Join(saveLocation, urlClean)
	log.Info("saving asset: %s : %s : %s", url, fullSaveLocation)

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
	} else {
		log.Info("successfully saved file")
	}

	metadata := fileData.CreateFileData(urlClean, fullSaveLocation, fileType)
	fileDataStore.StoreFileData(metadata)
}
