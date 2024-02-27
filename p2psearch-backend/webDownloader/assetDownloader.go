/* Package downloads assets required for the webpage */

package webDownloader

import (
	"fmt"
	"github.com/alexkefer/p2psearch-backend/fileData"
	"github.com/alexkefer/p2psearch-backend/fileTypes"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strings"
)

func downloadAllAssets(baseURL, htmlContent string, fileDataStore *fileData.FileDataStore) error {
	tokenizer := html.NewTokenizer(strings.NewReader(htmlContent))
	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return nil
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			switch token.Data {
			case "link":
				for _, attr := range token.Attr {
					if attr.Key == "rel" && strings.Contains(attr.Val, "stylesheet") {
						if href, ok := getAttributeValue(token, "href"); ok {
							cssURL := buildURL(baseURL, href)
							DownloadCSS(cssURL, fileDataStore)
						}
					}
				}
			case "script":
				for _, attr := range token.Attr {
					if attr.Key == "src" {
						if src, ok := getAttributeValue(token, "src"); ok {
							jsURL := buildURL(baseURL, src)
							DownloadJS(jsURL, fileDataStore)
						}
					}
				}
			case "img", "audio", "video":
				for _, attr := range token.Attr {
					if attr.Key == "src" {
						if src, ok := getAttributeValue(token, "src"); ok {
							println("src: " + src)
							downloadAsset(baseURL, src, fileDataStore)
						}
					}
				}
			}
		}
	}
}

// Downloads the css file given the url
func DownloadCSS(url string, fileDataStore *fileData.FileDataStore) {
	// takes in url and returns the css file
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
	SaveFile(data, CleanUrl(url), fileTypes.Css, fileDataStore)
}

// Downloads required js files
func DownloadJS(url string, fileDataStore *fileData.FileDataStore) {
	// takes in url and returns the js file
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("error getting url:" + string(rune(resp.StatusCode)))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	SaveFile(data, CleanUrl(url), fileTypes.Javascript, fileDataStore)
}

// Downloads various assets given the url
func downloadAsset(baseURL, url string, fileDataStore *fileData.FileDataStore) {
	println("downloading asset: " + url + " " + baseURL)
	resp, err := http.Get(buildURL(baseURL, url))
	if err != nil {
		println("Cannot access asset: " + buildURL(baseURL, url) + " " + err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		if resp.StatusCode == 404 {
			fmt.Println("Asset not found " + buildURL(baseURL, url))
		}
		println("error getting url:" + string(rune(resp.StatusCode)))
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	SaveFile(data, CleanUrl(url), "", fileDataStore)
}

/* Helper functions */

func getAttributeValue(token html.Token, key string) (string, bool) {
	for _, attr := range token.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}
