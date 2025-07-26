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

var urls = []string{"https://go.dev", "https://golang.org"}

var sFlag = flag.String("s", "", "keywords to search for IN TITLE")

var docs = []crawler.Document{}

func main() {
	flag.Parse()
	if *sFlag == "" {
		fmt.Println("Exit: Target word not set (use -s argument to set target word)")
		return
	}
	*sFlag = strings.TrimSpace(*sFlag)
	*sFlag = strings.ToLower(*sFlag)

	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		scanAll()
	} else {
		loadAll()
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

func scanAll() {
	for _, u := range urls {
		scan(u)
	}

	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0)
	if err != nil {
		fmt.Printf("Error: %v during opening file: %s\n", err, filepath)
	}

	save(f)

	err = f.Close()
	if err != nil {
		fmt.Printf("Error: %v during closing file: %s\n", err, filepath)
	}

}

func loadAll() {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error: %v during opening file: %s\n", err, filepath)
	}

	load(f)

	err = f.Close()
	if err != nil {
		fmt.Printf("Error: %v during closing file: %s\n", err, filepath)
	}
}

func load(r io.Reader) {
	dec := gob.NewDecoder(r)
	err := dec.Decode(&docs)
	if err != nil {
		fmt.Printf("Failed to load docs\t>!!< Error:%v\n", err)
	}

	for _, d := range docs {
		index.Add(d)
	}
}

func save(w io.Writer) {
	enc := gob.NewEncoder(w)
	err := enc.Encode(docs)
	if err != nil {
		fmt.Printf("Failed to save docs\t>!!< Error:%v\n", err)
	}
}
