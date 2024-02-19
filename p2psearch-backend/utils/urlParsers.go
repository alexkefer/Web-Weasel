/* Utility functions to parse various file information out of the URL */

package utils

import "strings"

func UrlCleaner(url string) string {
	// takes in url and returns the cleaned url (removes http(s):// and www.)
	if len(url) >= 8 && url[:8] == "https://" {
		url = url[8:]
	}
	if len(url) >= 7 && url[:7] == "http://" {
		url = url[7:]
	}
	if len(url) >= 4 && url[:4] == "www." {
		url = url[4:]
	}
	if url[len(url)-1] == '/' {
		url = url[:len(url)-1]
	}
	if url[0] == '/' {
		url = url[1:]
	}
	return url
}

func ParseSourceLocation(url string) string {
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

func ParsePageLocation(url string) string {
	// takes in url and returns the location of the page
	for i := len(url) - 1; i >= 0; i-- {
		if url[i] == '/' {
			url = url[:i]
			break
		}
	}
	return "savedPages/" + UrlCleaner(url)
}

func ParsePageName(url string) string {
	// takes in url and returns the name of the page
	for i := len(url) - 1; i >= 0; i-- {
		if url[i] == '/' {
			url = url[i+1:]
			break
		}
	}
	return url
}

func BuildURL(baseURL, assetURL string) string {
	// takes in a base url and an asset url and returns the full url
	if strings.HasPrefix(assetURL, "http") || strings.HasPrefix(assetURL, "https") {
		return assetURL
	} else if strings.HasPrefix(assetURL, "/") || strings.HasSuffix(baseURL, "/") {
		return baseURL + assetURL
	} else if strings.HasPrefix(assetURL, "//") {
		return "https:" + assetURL
	} else {
		return baseURL + "/" + UrlCleaner(assetURL)
	}
}
