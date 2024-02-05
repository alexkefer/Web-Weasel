/* Function and helpers to save pages to file at specified location*/

package webDownloader

import (
	"fmt"
	"os"
)

func savePage(context string, url string, saveLocation string, fileType string) {
	err := os.MkdirAll(parsePageLocation(url), os.ModePerm)
	println("Saving Asset: " + url + " : " + urlCleaner(url) + " : " + buildURL(saveLocation, url))
	if err != nil {
		fmt.Println(err)
		return
	}
	// takes in the context of the page and saves it to the save location
	file, err := os.OpenFile(saveLocation+"/"+urlCleaner(url)+fileType, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err2 := file.WriteString(context)
	if err2 != nil {
		fmt.Println("Error writing to file")
	} else {
		fmt.Println("Successfully saved file")
	}
}
