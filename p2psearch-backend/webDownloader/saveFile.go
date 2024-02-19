/* Function and helpers to save pages to file at specified location*/

package webDownloader

import (
	"github.com/alexkefer/p2psearch-backend/log"
	"os"
)

func savePage(context string, url string, saveLocation string, fileType string) {
	err := os.MkdirAll(parsePageLocation(url), os.ModePerm)
	log.Info("saving asset: %s : %s : %s", url, urlCleaner(url), buildURL(saveLocation, url))
	if err != nil {
		log.Error("error saving asset: %s", err)
		return
	}
	// takes in the context of the page and saves it to the save location
	file, err := os.OpenFile(saveLocation+"/"+urlCleaner(url)+fileType, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Error("error opening file: %s", err)
		return
	}
	defer file.Close()

	_, err2 := file.WriteString(context)
	if err2 != nil {
		log.Error("error writing to file: %s", err2)
	} else {
		log.Info("successfully saved file")
	}
}
