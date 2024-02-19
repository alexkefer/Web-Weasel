// Alex Kefer // January 2023 // Package to download webpages to be able to run them locally
// Will include options for all pages or just the page itself and where to save it
// Helper functions will assist in translating the html css and js files

package webDownloader

import (
	"errors"
	"github.com/alexkefer/p2psearch-backend/fileData"
	"github.com/alexkefer/p2psearch-backend/fileTypes"
	"github.com/alexkefer/p2psearch-backend/log"
	"github.com/alexkefer/p2psearch-backend/utils"
	"io"
	"net/http"
)

func BuildDownloadedWebpage(url string, fileDataStore *fileData.FileDataStore) error {
	pageHtml, err := DownloadPage(url)
	if err != nil {
		log.Warn("error downloading page: %s", err)
		return err
	}
	err2 := downloadAllAssets(utils.ParseSourceLocation(url), pageHtml, fileDataStore)
	if err2 != nil {
		log.Warn("error downloading assets: %s", err2)
		return err2
	}
	pageHtml = regexHtml(pageHtml, url)
	savePage([]byte(pageHtml), url, fileTypes.Html, fileDataStore)
	log.Info("downloaded webpage at %s", url)
	return nil
}

func DownloadPage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Warn("couldn't download web page at %s", url)
		return "", errors.New("error getting url: " + url)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	content := string(data)
	// Print the content
	// fmt.Println(content)
	return content, nil
}
