package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"go_search/05-homework/pkg/crawler"
	"go_search/05-homework/pkg/crawler/spider"
	"go_search/05-homework/pkg/index"
	"io"
	"math/rand/v2"
	"os"
	"slices"
	"strings"
)

const depth = 2

const filepath = "docs"

func main() {
	var docs []crawler.Document
	urls := []string{"https://go.dev", "https://golang.org"}
	sFlag := flag.String("s", "", "keywords to search for IN TITLE")

	flag.Parse()
	if *sFlag == "" {
		fmt.Println("Exit: Target word not set (use -s argument to set target word)")
		return
	}
	*sFlag = strings.TrimSpace(*sFlag)
	*sFlag = strings.ToLower(*sFlag)

	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		docs = scanAll(urls)
	} else {
		f, err := os.Open(filepath)
		if err != nil {
			fmt.Printf("Error: %v during opening file: %s\n", err, filepath)
		}

		docs = load(f)

		err = f.Close()
		if err != nil {
			fmt.Printf("Error: %v during closing file: %s\n", err, filepath)
		}
	}

	fmt.Printf("Searching in Index for phrase: %s\n", *sFlag)
	res := index.Search(*sFlag)
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

func scan(url string) []crawler.Document {
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
	slices.SortFunc(res, func(a crawler.Document, b crawler.Document) int {
		return a.ID - b.ID
	})
	return res
}

func scanAll(urls []string) []crawler.Document {
	res := make([]crawler.Document, 0)
	for _, u := range urls {
		res = append(res, scan(u)...)
	}

	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0)
	if err != nil {
		fmt.Printf("Error: %v during opening file: %s\n", err, filepath)
	}

	save(f, res)

	err = f.Close()
	if err != nil {
		fmt.Printf("Error: %v during closing file: %s\n", err, filepath)
	}

	return res
}

func load(r io.Reader) []crawler.Document {
	dec := gob.NewDecoder(r)
	res := make([]crawler.Document, 0)
	err := dec.Decode(&res)
	if err != nil {
		fmt.Printf("Failed to load docs\t>!!< Error:%v\n", err)
	}

	for _, d := range res {
		index.Add(d)
	}

	return res
}

func save(w io.Writer, d []crawler.Document) {
	enc := gob.NewEncoder(w)
	err := enc.Encode(d)
	if err != nil {
		fmt.Printf("Failed to save docs\t>!!< Error:%v\n", err)
	}
}
