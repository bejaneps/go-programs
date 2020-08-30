package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(urlToFQDN("https://twitter.com"))
	fmt.Println(urlToFQDN("http://twitter.com"))
	fmt.Println(urlToFQDN("https://www.twitter.com"))
	fmt.Println(urlToFQDN("https://facebook.com"))
	fmt.Println(urlToFQDN("http://twitter.com"))
	fmt.Println(urlToFQDN("https://www.twitter.com"))
}

func urlToFQDN(url string) string {
	if strings.HasPrefix(url, "0.0.0.0 ") {
		url = strings.Trim(url, "0.0.0.0")
	} else if strings.HasPrefix(url, "::") {
		url = strings.Trim(url, "::")
	} else if strings.HasPrefix(url, "http://") {
		url = strings.Trim(url, "http:")
	} else if strings.HasPrefix(url, "https://") {
		url = strings.Trim(url, "https:")
	}
	url = strings.Replace(url, " ", "", -1) // remove all spaces from url
	url = strings.Replace(url, "/", "", -1) // remove all url path separators

	// convert url to fqdn
	return url + "."
}
