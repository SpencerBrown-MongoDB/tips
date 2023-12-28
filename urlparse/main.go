package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	thisURL := "12345678"
	parsedURL, err := url.Parse(thisURL)
	elements := strings.Split(parsedURL.Path, "/")
	fmt.Printf("URL %s; val %s; err %v\n", thisURL, parsedURL.Path, err)
	for i, element := range elements {
		fmt.Printf("%2d %s\n", i, element)
	}
	realURL := "https://cloud.exampledb.com/v2/64e75cdd2238f41d106e2dd8#/overview"
	escapedURL := url.QueryEscape(realURL)
	realURL2, err := url.QueryUnescape(escapedURL)
	if err != nil {
		fmt.Printf("Unescape error: %v\n", err)
	}
	fmt.Printf("URL: %s\nEscaped: %s\nUnescaped: %s\n", realURL, escapedURL, realURL2)
}
