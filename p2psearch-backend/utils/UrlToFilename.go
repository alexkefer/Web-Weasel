package utils

import "strings"

// UrlToFilename trims each filename to a basic name and replaces slashes with dashes.
func UrlToFilename(url string) string {
	url = strings.TrimPrefix(url, "https://")
	url = strings.TrimPrefix(url, "http://")
	url = strings.ReplaceAll(url, "/", "_")
	return url
}
