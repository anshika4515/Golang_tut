package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://elegant-sunburst-72a415.netlify.app/"

func main() {
	fmt.Println("About handling urls")
	fmt.Println(myurl)
	//parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port)
	fmt.Println(result.RawQuery)

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user-hitesh",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
