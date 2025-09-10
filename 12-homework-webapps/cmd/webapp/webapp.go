package main

import (
	"encoding/json"
	"fmt"
	"go_search/12-homework-webapps/pkg/index"
	"go_search/12-homework-webapps/pkg/scanner"
	"log"
	"net/http"
)

const depth = 2

func main() {

	urls := []string{"https://go.dev", "https://golang.org"}

	echan := make(chan error, 10)
	go scanner.ScanAll(urls, depth, echan)
	go func() {
		err := <-echan
		fmt.Println(err)
	}()

	http.HandleFunc("/index", showIndex)
	http.HandleFunc("/docs", showDocs)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func showIndex(w http.ResponseWriter, r *http.Request) {
	jIndex, err := json.MarshalIndent(index.Index, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}

	_, err = w.Write(jIndex)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}
}

func showDocs(w http.ResponseWriter, r *http.Request) {
	jDocs, err := json.MarshalIndent(scanner.Docs, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}

	_, err = w.Write(jDocs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}
}
