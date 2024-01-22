// Alex Kefer // January 2023 // Package to download webpages to be able to run them locally
// Will include options for all pages or just the page itself and where to save it
// Helper functions will assist in translating the html css and js files

package webDownloader

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"time"
)

func DownloadPage(url string) {
	// Create a new context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Navigate to the URL
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
	); err != nil {
		log.Fatal(err)
	}

	// Wait for some seconds to let the JavaScript execute
	time.Sleep(2 * time.Second)

	// Get the page content
	var content string
	if err := chromedp.Run(ctx,
		chromedp.OuterHTML("html", &content),
	); err != nil {
		log.Fatal(err)
	}

	// Print the content
	fmt.Println(content)
	savePage(content, parseUrl(url), "savedPages")
}

func parseUrl(url string) string {
	// takes in a url and returns the parsed url
	if url[:8] == "https://" {
		url = url[8:]
	}
	println(url)
	return url
}

func savePage(context string, url string, saveLocation string) {
	// takes in the context of the page and saves it to the save location
	file, err := os.OpenFile(saveLocation+"/"+url, os.O_RDWR|os.O_CREATE, 0644)
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
