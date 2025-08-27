package main

import (
	"go_search/12-homework/pkg/crawler"
	"go_search/12-homework/pkg/index"
	"go_search/12-homework/pkg/scanner"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var testMux *http.ServeMux

func TestMain(m *testing.M) {
	testMux = http.NewServeMux()
	testMux.HandleFunc("/index", showIndex)
	testMux.HandleFunc("/docs", showDocs)
	m.Run()
}

func Test_showIndex(t *testing.T) {
	index.Index = map[string][]int{"go": {123, 321}, "lang": {444}}

	req := httptest.NewRequest(http.MethodGet, "/index", nil)

	rr := httptest.NewRecorder()

	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, хотели %d", rr.Code, http.StatusOK)
	}

	body := rr.Body.String()
	if !strings.Contains(body, "go") || !strings.Contains(body, "321") || !strings.Contains(body, "lang") {
		t.Fatal(body)
	}
}

func Test_showDocs(t *testing.T) {
	scanner.Docs = []crawler.Document{
		{
			ID:    444,
			URL:   "http://www.testurl.com",
			Title: "Title",
			Body:  "Body",
		},
		{
			ID:    555,
			URL:   "http://www.test2url.com",
			Title: "Title2",
			Body:  "Body2",
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/docs", nil)

	rr := httptest.NewRecorder()

	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, хотели %d", rr.Code, http.StatusOK)
	}

	body := rr.Body.String()
	if !strings.Contains(body, "555") || !strings.Contains(body, "test2") || !strings.Contains(body, "url.com") {
		t.Fatal(body)
	}
}
