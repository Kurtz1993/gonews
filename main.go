package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type SitemapIndex struct {
	// Locations property of type list of Location which is inside the xml <sitemap> tag
	Locations	[]Location	`xml:"sitemap"`
}

type Location struct {
	// Loc property of type string which is inside the xml <loc> tag
	Loc	string	`xml:"loc"`
	// LastModified property of type string which is inside the xml <lastmod> tag
	LastModified	string	`xml:"lastmod"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	response, _ := http.Get("https://www.nytimes.com/sitemaps/new/cooking.xml.gz")
	bytes, _ := io.ReadAll(response.Body)
	response.Body.Close()

	var s SitemapIndex

	// Parses the bytes from the XML response to the specified type
	// of the s variable
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations)

	// This works similar to a for-of loop. The "_" will be the index
	// location represents the value
	for _, location := range(s.Locations) {
		fmt.Printf("\n%s", location)
	}
}
