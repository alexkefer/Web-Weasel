/* This is a helper utility built to regex through the html and modify the locations to where they are downloaded rather than their links */

package webDownloader

import "regexp"

func regexHtml(html string, baseURL string) string {
	// regex through the html and modify the locations to where they are downloaded rather than their links
	//Expression Pattern
	cssURLs := regexp.MustCompile(`url\(['"]?(.*?)['"]?\)`)

	// Used regular expression to find all the css links
	modifiedHtml := cssURLs.ReplaceAllStringFunc(html, func(link string) string {
		// Extract URL from the matched string
		cssURL := cssURLs.FindStringSubmatch(link)
		// If the URL is a relative path, append the base URL
		if len(cssURL) > 1 {
			url := cssURL[1]

			//Modify URL if need be
			updatedURL := buildURL(baseURL, url)

			return "url(" + updatedURL + ")"
		}
		return link
	})
	return modifiedHtml
}
