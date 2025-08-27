package main

import (
	"fmt"
	"go_search/12-homework/pkg/crawler"
	"go_search/12-homework/pkg/index"
	"go_search/12-homework/pkg/scanner"
	"slices"
)

const depth = 2

func main() {
	var docs []crawler.Document
	urls := []string{"https://go.dev", "https://golang.org"}

	var sreq string

	docs = scanner.ScanAll(urls, depth)

	fmt.Printf("Searching in Index for phrase: %s\n", sreq)
	res := index.Search(sreq)
	fmt.Printf("Match found: %d\n in documents #%v\n", len(res), res)

	for _, id := range res {
		i, ok := slices.BinarySearchFunc(docs, id, func(d crawler.Document, n int) int {
			return d.ID - n
		})
		if ok {
			fmt.Printf("#%d\t\t%s\t\t%s\n", docs[i].ID, docs[i].Title, docs[i].URL)
		}
	}
}
