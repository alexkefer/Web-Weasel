/* Keagan Edwards & Alex Kefer - This is a helper utility built to regex through the html and modify the locations to where they are downloaded rather than their links */

package webDownloader

import (
	"github.com/alexkefer/p2psearch-backend/fileData"
	"github.com/alexkefer/p2psearch-backend/log"
	"github.com/alexkefer/p2psearch-backend/utils"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strings"
)

// Sends a request to the URL in order to download the webpage.
func retrieveAsset(url string) ([]byte, string) {
	// takes in the url and returns the asset
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "P2PWebCache")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("error downloading asset:", url, err)
		return nil, ""
	}

	contentType := resp.Header.Get("Content-Type")

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("error reading asset content:", err)
		return nil, contentType
	}
	return body, contentType
}

/*
DownloadAllAssets tokenizes through the html content of the page
and finds key content needed to properly display the web page.
This can include CSS, JS and photos.
This can greatly be expanded upon to better handle advanced websites.
*/
func DownloadAllAssets(url, htmlContent string, fileStore *fileData.FileDataStore) string {
	tokenizer := html.NewTokenizer(strings.NewReader(htmlContent))
	modifiedHtml := ""
	log.Error("content url: %s", url)
	shortUrl := shortenUrl(url)

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return modifiedHtml
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			switch token.Data {
			case "link": // Download CSS
				for i, attr := range token.Attr {
					if attr.Key == "href" {
						rel, ok := getAttributeValue(token, "rel")
						if ok {
							if rel == "stylesheet" {
								link := shortUrl + attr.Val
								log.Debug("retrieving stylesheet asset: " + link)
								content, contentType := retrieveAsset(link)
								if content != nil {
									filename := utils.UrlToFilename(link)
									SaveFile(content, filename, contentType, fileStore)
									attr.Val = "/retrieve?path=" + filename
									token.Attr[i] = attr
								}
							} else if rel == "preload" {
								as, ok := getAttributeValue(token, "as")
								if ok && (as == "style") {
									link := shortUrl + attr.Val
									log.Debug("retrieving stylesheet asset: " + link)
									content, contentType := retrieveAsset(link)
									if content != nil {
										filename := utils.UrlToFilename(link)
										SaveFile(content, filename, contentType, fileStore)
										attr.Val = "/retrieve?path=" + filename
										token.Attr[i] = attr
									}
								}
							}
						}
					}
				}
			case "script": // Download JS
				for i, attr := range token.Attr {
					if attr.Key == "src" {
						link := shortUrl + attr.Val
						log.Debug("retrieving Asset: " + link)
						content, contentType := retrieveAsset(link)
						if content != nil {
							filename := utils.UrlToFilename(link)
							SaveFile(content, filename, contentType, fileStore)
							attr.Val = "/retrieve?path=" + filename
							token.Attr[i] = attr
						}
					}
				}
			case "img": // Download Images
				for i, attr := range token.Attr {
					if attr.Key == "src" {
						link := shortUrl + attr.Val
						log.Debug("retrieving Asset: " + link)
						content, contentType := retrieveAsset(link)
						if content != nil {
							filename := utils.UrlToFilename(link)
							SaveFile(content, filename, contentType, fileStore)
							attr.Val = "/retrieve?path=" + filename
							token.Attr[i] = attr
						}
					}
				}
			}

			modifiedHtml += html.UnescapeString(token.String())

		default:
			token := tokenizer.Token()
			modifiedHtml += html.UnescapeString(token.String())
		}
	}
}

/* Helper Functions */

// Returns attribute value for a given key from an html token
func getAttributeValue(token html.Token, key string) (string, bool) {
	for _, attr := range token.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

// Shortens the URL by removing the prefix for ease of use in searching for a URL.
func shortenUrl(fullUrl string) string {
	cutUrl, cut := strings.CutPrefix(fullUrl, "https://")
	if cut {
		front, _, found := strings.Cut(cutUrl, "/")
		if found {
			return "https://" + front + "/"
		}
	}

	cutUrl, cut = strings.CutPrefix(fullUrl, "http://")
	if cut {
		front, _, found := strings.Cut(cutUrl, "/")
		if found {
			return "https://" + front + "/"
		}
	}

	return fullUrl
}
