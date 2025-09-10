package scanner

import (
	"go_search/12-homework-webapps/pkg/crawler"
	"go_search/12-homework-webapps/pkg/crawler/spider"
	"go_search/12-homework-webapps/pkg/index"
	"math/rand/v2"
	"slices"
)

var Docs []crawler.Document

func Scan(url string, depth int) ([]crawler.Document, error) {
	res := make([]crawler.Document, 0)
	spd := spider.New()
	scans, err := spd.Scan(url, depth)
	if err != nil {
		return nil, err
	}

	for _, s := range scans {
		r := crawler.Document{ID: int(rand.Float32() * 10000), Title: s.Title, URL: s.URL}
		index.Add(r)
		res = append(res, r)
	}

	return res, nil
}

func ScanAll(urls []string, d int, echan chan error) {
	for _, u := range urls {
		res, err := Scan(u, d)
		if err != nil {
			echan <- err
			continue
		}
		Docs = append(Docs, res...)
	}

	slices.SortFunc(Docs, func(a crawler.Document, b crawler.Document) int {
		return a.ID - b.ID
	})
	close(echan)
}
