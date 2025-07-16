package main

import (
	"flag"
	"fmt"
	"go_search/03-homework/pkg/crawler"
	"go_search/03-homework/pkg/crawler/spider"
	"go_search/03-homework/pkg/index"
	"math/rand/v2"
	"slices"
	"strings"
)

const depth = 2

var urls = []string{"https://go.dev", "https://golang.org"}

var sFlag = flag.String("s", "", "keywords to search for IN TITLE")

var docs = []crawler.Document{}

func main() {
	flag.Parse()
	*sFlag = strings.TrimSpace(*sFlag)
	*sFlag = strings.ToLower(*sFlag)

	for _, u := range urls {
		scan(u)
	}

	if *sFlag == "" {
		return
	}

	fmt.Printf("Searching in Index for phrase: %s\n", *sFlag)
	res := index.Search(*sFlag)
	fmt.Printf("Match found in documents #%v\n", res)

	for _, id := range res {
		i, ok := slices.BinarySearchFunc(docs, id, func(d crawler.Document, n int) int {
			return d.ID - n
		})
		if ok {
			fmt.Printf("#%d\t\t%s\t\t%s\n", docs[i].ID, docs[i].Title, docs[i].URL)
		}
	}
}

func scan(url string) {
	spd := spider.New()
	scans, err := spd.Scan(url, depth)
	if err != nil {
		fmt.Printf("Scan failed for url:%s\n\tError:%v", url, err)
		return
	}
	for _, s := range scans {
		doc := crawler.Document{ID: int(rand.Float32() * 10000), Title: s.Title, URL: s.URL}
		index.Add(doc)
		docs = append(docs, doc)
	}
	slices.SortFunc(docs, func(a crawler.Document, b crawler.Document) int {
		return a.ID - b.ID
	})
}
