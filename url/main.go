package main

import (
	"fmt"
	"net/url"
)
func main() {
	fmt.Println("Learning Url")
	myUrl := "https://www.youtube.com/watch?v=vu6ZQ-t1sUk&list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&index=26"

	parsedURL, _ := url.Parse(myUrl)

	fmt.Printf("Type of URL: %T\n", parsedURL)
	fmt.Println("Scheme", parsedURL.Scheme)
	fmt.Println("RawQurey", parsedURL.RawQuery)
	fmt.Println("Path", parsedURL.Path)

}