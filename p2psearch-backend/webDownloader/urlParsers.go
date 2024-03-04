/* Utility functions to parse various file information out of the URL */

package webDownloader

import "strings"

func CleanUrl(url string) string {
	url = urlCleaner(url)
	url = strings.ReplaceAll(url, "/", "_")
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

func buildURL(baseURL, assetURL string) string {
	// takes in a base url and an asset url and returns the full url
	if strings.HasPrefix(assetURL, "http") || strings.HasPrefix(assetURL, "https") {
		return assetURL
	} else if strings.HasPrefix(assetURL, "/") || strings.HasSuffix(baseURL, "/") {
		return baseURL + assetURL
	} else if strings.HasPrefix(assetURL, "//") {
		return "https:" + assetURL
	} else {
		return baseURL + "/" + urlCleaner(assetURL)
	}
}
