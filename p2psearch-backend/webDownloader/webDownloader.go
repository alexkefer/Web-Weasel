// Alex Kefer // January 2023 // Package to download webpages to be able to run them locally
// Will include options for all pages or just the page itself and where to save it
// Helper functions will assist in translating the html css and js files

package webDownloader

import (
	"fmt"
	"io"
	"net/http"
)

func BuildDownloadedWebpage(url string) {
	pageHtml, err := DownloadPage(url)
	if err != nil {
		fmt.Println("error downloading page:", err)
		return
	}
	err2 := downloadAllAssets(parseSourceLocation(url), pageHtml)
	if err2 != nil {
		fmt.Println("error downloading assets:", err2)
		return
	}
	pageHtml = regexHtml(pageHtml, url, "savedPages/")
	savePage(pageHtml, url, "savedPages", ".html")
	println("Successfully downloaded webpage: " + url)
}

func DownloadPage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("error getting url: " + url)
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
