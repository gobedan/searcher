package scanner

import (
	"fmt"
	"go_search/12-homework/pkg/crawler"
	"go_search/12-homework/pkg/crawler/spider"
	"go_search/12-homework/pkg/index"
	"math/rand/v2"
	"slices"
)

var Docs []crawler.Document

func Scan(url string, depth int) []crawler.Document {
	res := make([]crawler.Document, 0)
	spd := spider.New()
	scans, err := spd.Scan(url, depth)
	if err != nil {
		fmt.Printf("Scan failed for url:%s\n\tError:%v", url, err)
		return res
	}

	for _, s := range scans {
		r := crawler.Document{ID: int(rand.Float32() * 10000), Title: s.Title, URL: s.URL}
		index.Add(r)
		res = append(res, r)
	}

	return res
}

func ScanAll(urls []string, d int) {
	for _, u := range urls {
		Docs = append(Docs, Scan(u, d)...)
	}

	slices.SortFunc(Docs, func(a crawler.Document, b crawler.Document) int {
		return a.ID - b.ID
	})
}
