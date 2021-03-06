package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"sync"
	"text/template"
	"time"
)

var wg sync.WaitGroup

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

type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
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

	p := NewsAggPage{Title: "Amazing News Aggregator", News: newsMap}
	t, _ := template.ParseFiles("news.gohtml")

	fmt.Println(t.Execute(w, p))
}

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	// Used to notify the wait group that this routine is done
	wg.Done()
}

func main() {
	//http.HandleFunc("/", newsAggHandler)

	// http.ListenAndServe(":5000", nil)

	// Add 1 to the wait group before runing our goroutine
	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("There")

	// Wait until all goroutines are done
	wg.Wait()
}
