//  Keagan Edwards & Alex Kefer - January 2023 - Package to download webpages to be able to run them locally

package webDownloader

import (
	"errors"
	"github.com/alexkefer/p2psearch-backend/fileData"
	"github.com/alexkefer/p2psearch-backend/fileTypes"
	"github.com/alexkefer/p2psearch-backend/log"
	"github.com/alexkefer/p2psearch-backend/utils"
	"io"
	"net/http"
	"strings"
)

/*
CacheResource first uses download Resource to read the html page.
Then uses the UrlToFilename function to build a directory for the files.
Afterward it proceeds to sort by its content and download all assets including its base html file.
*/
func CacheResource(url string, fileDataStore *fileData.FileDataStore) error {
	content, contentType, err := downloadResource(url)
	if err != nil {
		log.Warn("error downloading page: %s", err)
		return err
	}

	filename := utils.UrlToFilename(url)

	if strings.Contains(contentType, "text/html") {
		modifiedHtml := DownloadAllAssets(url, string(content), fileDataStore)
		SaveFile([]byte(modifiedHtml), filename, fileTypes.Html, fileDataStore)
		log.Info("cached webpage at %s", url)
	} else {
		SaveFile(content, filename, contentType, fileDataStore)
		log.Info("cached resource at %s", url)
	}

	return nil
}

// Download Resource reads the given webpage from the url string.
func downloadResource(url string) ([]byte, string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, "", err
	}

	if resp.StatusCode != 200 {
		log.Warn("couldn't download web page at %s", url)
		return nil, "", errors.New("error getting url: " + url)
	}

	contentType := resp.Header.Get("Content-Type")
	data, err := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		return nil, contentType, errors.New("error getting url: " + url)
	}
	// Print the content
	// fmt.Println(content)
	return data, contentType, nil
}
