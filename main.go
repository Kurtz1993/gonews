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
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

func main() {
	var n News
	newsMap := make(map[string]NewsMap)

	// location represents the value
	// for _, location := range s.Locations {
	response, _ := http.Get("https://www.nytimes.com/sitemaps/new/news.xml.gz")
	bytes, _ := io.ReadAll(response.Body)
	response.Body.Close()

	xml.Unmarshal(bytes, &n)

	// This works similar to a for-of loop. The "_" will be the value.
	// It can also be simplified by removing the "_"
	for idx, _ := range n.Keywords {
		newsMap[n.Titles[idx]] = NewsMap{
			Keyword:  n.Keywords[idx],
			Location: n.Locations[idx],
		}
	}
	// }

	for idx, data := range newsMap {
		fmt.Println("\n\n\n", "Title:", idx)
		fmt.Println("\n", "Keyword:", data.Keyword)
		fmt.Println("\n", "Location:", data.Location)
	}
}
