// Alex Kefer // January 2023 // Package to download webpages to be able to run them locally
// Will include options for all pages or just the page itself and where to save it
// Helper functions will assist in translating the html css and js files

package webDownloader

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"strings"
)

func BuildDownloadedWebpage(url string) {
	pageHtml, err := DownloadPage(url)
	if err != nil {
		fmt.Println("error downloading page:", err)
		return
	}
	savePage(pageHtml, url, "savedPages", ".html")
	err2 := downloadAllAssets(parseSourceLocation(url), pageHtml)
	if err2 != nil {
		fmt.Println("error downloading assets:", err2)
		return
	}
	println("Successfully downloaded webpage: " + url)
}

func DownloadPage(url string) (string, error) {
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
	content := string(data)
	// Print the content
	fmt.Println(content)

	return content, nil
}

func DownloadCSS(url string) {
	// takes in url and returns the css file
	print(url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	println(url)
	if resp.StatusCode != 200 {
		panic("error getting url:" + string(rune(resp.StatusCode)))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	savePage(string(data), url, "savedPages", ".css")
}

func downloadAllAssets(baseURL, htmlContent string) error {
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
							DownloadCSS(cssURL)
						}
					}
				}
			case "script":
				for _, attr := range token.Attr {
					if attr.Key == "src" {
						if src, ok := getAttributeValue(token, "src"); ok {
							jsURL := buildURL(baseURL, src)
							DownloadJS(jsURL)
						}
					}
				}
			case "img", "audio", "video":
				for _, attr := range token.Attr {
					if attr.Key == "src" {
						if src, ok := getAttributeValue(token, "src"); ok {
							println("src: " + src)
							downloadAsset(baseURL, src)
						}
					}
				}
			}
		}
	}
}

func DownloadJS(url string) {
	// takes in url and returns the js file
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	println(url)
	if resp.StatusCode != 200 {
		panic("error getting url:" + string(rune(resp.StatusCode)))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	savePage(string(data), url, "savedPages", ".js")
}

func getAttributeValue(token html.Token, key string) (string, bool) {
	for _, attr := range token.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func buildURL(baseURL, assetURL string) string {
	// takes in a base url and an asset url and returns the full url
	if strings.HasPrefix(assetURL, "http") || strings.HasPrefix(assetURL, "https") {
		return assetURL
	} else if strings.HasPrefix(assetURL, "/") || strings.HasSuffix(baseURL, "/") {
		return baseURL + assetURL
	} else if strings.HasPrefix(assetURL, "//") {
		println("B: " + "https:" + assetURL)
		return "https:" + assetURL
	} else {
		println("A: " + baseURL + "/" + urlCleaner(assetURL))
		return baseURL + "/" + urlCleaner(assetURL)
	}
}

// function struggles with external sites, idk why im tired. will fix later
func downloadAsset(baseURL, url string) {
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
	content := string(data)
	savePage(content, url, "savedPages", "")
}

func urlCleaner(url string) string {
	// takes in url and returns the cleaned url (removes http(s):// and www.)
	println(url)
	if len(url) >= 8 && url[:8] == "https://" {
		url = url[8:]
	}
	if len(url) >= 7 && url[:7] == "http://" {
		url = url[7:]
	}
	if len(url) >= 4 && url[:4] == "www." {
		url = url[4:]
	}
	return url
}

func parseSourceLocation(url string) string {
	// takes in url and returns the location of the source website (for assets)
	i := 0
	if len(url) >= 8 && url[:8] == "https://" {
		i = 8
	}
	if len(url) >= 7 && url[:7] == "http://" {
		i = 7
	}
	for ; i < len(url); i++ {
		if url[i] == '/' {
			url = url[:i]
			break
		}
	}
	return url
}

func parsePageLocation(url string) string {
	// takes in url and returns the location of the page
	for i := len(url) - 1; i >= 0; i-- {
		if url[i] == '/' {
			url = url[:i]
			break
		}
	}
	return "savedPages/" + urlCleaner(url)
}

func parsePageName(url string) string {
	// takes in url and returns the name of the page
	for i := len(url) - 1; i >= 0; i-- {
		if url[i] == '/' {
			url = url[i+1:]
			break
		}
	}
	return url
}

func savePage(context string, url string, saveLocation string, fileType string) {
	err := os.MkdirAll(parsePageLocation(url), os.ModePerm)
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
