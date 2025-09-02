package main

import (
	"go_search/13-homework/pkg/crawler"
	"go_search/13-homework/pkg/index"
	"go_search/13-homework/pkg/scanner"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var testMux *http.ServeMux

func TestMain(m *testing.M) {
	/* 	testMux = http.NewServeMux()
	   	testMux.HandleFunc("/index", showIndex)
	   	testMux.HandleFunc("/docs", showDocs) */

	scanner.Docs = []crawler.Document{
		{
			ID:    444,
			URL:   "http://www.testurl.com",
			Title: "Title First",
			Body:  "Body One",
		},
		{
			ID:    555,
			URL:   "http://www.test2url.com",
			Title: "Title Second",
			Body:  "Body Two",
		},
	}

	index.Add(scanner.Docs[0])
	index.Add(scanner.Docs[1])

	m.Run()
}

func Test_showIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/index", nil)

	rr := httptest.NewRecorder()

	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, хотели %d", rr.Code, http.StatusOK)
	}

	body := rr.Body.String()
	if !strings.Contains(body, "title") || !strings.Contains(body, "555") || !strings.Contains(body, "444") {
		t.Fatal(body)
	}
}

func Test_showDocs(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/docs", nil)

	rr := httptest.NewRecorder()

	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, хотели %d", rr.Code, http.StatusOK)
	}

	body := rr.Body.String()
	if !strings.Contains(body, "One") || !strings.Contains(body, "Two") || !strings.Contains(body, ".com") {
		t.Fatal(body)
	}
}
