package main

import (
	"fmt"
	"go_search/13-homework-api/pkg/api"
	"go_search/13-homework-api/pkg/scanner"
	"log"
	"net/http"
)

const depth = 2

type server struct {
	api *api.API
}

func main() {

	urls := []string{"https://go.dev", "https://golang.org"}

	echan := make(chan error, 10)
	go scanner.ScanAll(urls, depth, echan)
	go func() {
		err := <-echan
		fmt.Println(err)
	}()

	srv := server{api: api.New()}

	err := http.ListenAndServe(":8080", srv.api.Router)
	if err != nil {
		log.Fatal(err)
	}

}
