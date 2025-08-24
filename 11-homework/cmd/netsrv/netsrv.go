package main

import (
	"bufio"
	"fmt"
	"go_search/11-homework/pkg/crawler"
	"go_search/11-homework/pkg/index"
	"go_search/11-homework/pkg/scanner"
	"log"
	"net"
	"slices"
	"time"
)

const depth = 2

var docs []crawler.Document

func main() {

	urls := []string{"https://go.dev", "https://golang.org"}

	docs = scanner.ScanAll(urls, depth)
	fmt.Println("Indexing finished, opening port...")

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		fmt.Println("Conn established")

		go handler(conn)
	}

}

func handler(conn net.Conn) {
	defer conn.Close()
	defer fmt.Println("Conn closed")

	conn.SetDeadline(time.Now().Add(time.Second * 100))

	r := bufio.NewReader(conn)
	for {
		sreq, _, err := r.ReadLine()
		if err != nil {
			return
		}

		fmt.Printf("Incoming search query: %s | ", sreq)
		_, err = conn.Write([]byte(search(string(sreq))))
		if err != nil {
			return
		}

		conn.SetDeadline(time.Now().Add(time.Second * 100))
	}

}

func search(sreq string) string {
	res := index.Search(sreq)
	fmt.Printf("Match found: %d\n in documents #%v\n", len(res), res)

	r := ""
	for _, id := range res {
		i, ok := slices.BinarySearchFunc(docs, id, func(d crawler.Document, n int) int {
			return d.ID - n
		})
		if ok {
			r += fmt.Sprintf("#%d | %s | %s\n", docs[i].ID, docs[i].Title, docs[i].URL)
		}
	}
	return r
}
