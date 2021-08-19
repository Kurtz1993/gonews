package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type SitemapIndex struct {
	// Locations property of type slice of Location which is inside the xml <sitemap> > <loc> tag
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

func main() {
	var s SitemapIndex
	var n News

	response, _ := http.Get("https://www.nytimes.com/sitemaps/new/news.xml.gz")
	bytes, _ := io.ReadAll(response.Body)
	response.Body.Close()

	// Parses the bytes from the XML response to the specified type
	// of the s variable
	xml.Unmarshal(bytes, &s)

	// This works similar to a for-of loop. The "_" will be the index
	// location represents the value
	for _, location := range s.Locations {
		response, _ := http.Get(location)
		bytes, _ := io.ReadAll(response.Body)
		response.Body.Close()

		xml.Unmarshal(bytes, &n)

		fmt.Println(n.Locations)
	}
}
