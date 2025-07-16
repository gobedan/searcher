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

var Docs = []crawler.Document{}

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

	res := search(*sFlag)

	for _, id := range res {
		i, ok := slices.BinarySearchFunc(Docs, id, func(d crawler.Document, n int) int {
			return d.ID - n
		})
		if ok {
			fmt.Printf("#%d\t\t%s\t\t%s\n", Docs[i].ID, Docs[i].Title, Docs[i].URL)
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
		Docs = append(Docs, doc)
	}
	slices.SortFunc(Docs, func(a crawler.Document, b crawler.Document) int {
		return a.ID - b.ID
	})
}

func search(s string) []int {
	res := []int{}
	fmt.Printf("Searching in Index for phrase: %s\n", *sFlag)
	for w, id := range index.Index {
		if w == s {
			res = append(res, id...)
			fmt.Printf("Match found in document #%d\n", id)
		}
	}
	return res
}
