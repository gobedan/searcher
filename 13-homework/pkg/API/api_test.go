package api

import (
	"encoding/json"
	"fmt"
	"go_search/13-homework/pkg/crawler"
	"go_search/13-homework/pkg/index"
	"go_search/13-homework/pkg/scanner"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
)

var api *API

func TestMain(m *testing.M) {
	api = New()
	scanner.Docs = []crawler.Document{
		{
			ID:    444,
			URL:   "http://www.test1url.com",
			Title: "Title First",
		},
		{
			ID:    555,
			URL:   "http://www.test2url.com",
			Title: "Title Second",
		},
	}

	index.Add(scanner.Docs[0])
	index.Add(scanner.Docs[1])

	m.Run()
}

func Test_Search(t *testing.T) {
	word := "First"
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/search?q=%s", word), nil)
	rr := httptest.NewRecorder()

	api.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, хотели %d", rr.Code, http.StatusOK)
	}

	got := rr.Body.String()
	want := fmt.Sprintf(`<a href="%s">%s</a><br>`, scanner.Docs[0].URL, scanner.Docs[0].Title)
	if got != want {
		t.Errorf("\n получили: \n%s\n хотели: \n%s\n", got, want)
	}
}

func Test_ShowDocs(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/docs", nil)
	rr := httptest.NewRecorder()

	api.Router.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, хотели %d", rr.Code, http.StatusOK)
	}

	got := rr.Body.String()
	want, _ := json.MarshalIndent(scanner.Docs, "", "\t")
	if got != string(want) {
		t.Errorf("\n получили: \n%s\n хотели: \n%s\n", got, want)
	}
}

func Test_AddDoc(t *testing.T) {
	title := "Added"
	url := "http://newdoc.com"
	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/docs?title=%s&url=%s", title, url), nil)
	rr := httptest.NewRecorder()

	api.Router.ServeHTTP(rr, req)
	if rr.Code != http.StatusCreated {
		t.Errorf("код неверен: получили %d, хотели %d", rr.Code, http.StatusCreated)
	}

	got := rr.Body.String()
	want := "Document successfully created!"
	if got != want {
		t.Errorf("\n получили: \n%s\n хотели: \n%s\n", got, want)
	}

	added := slices.ContainsFunc(scanner.Docs, func(doc crawler.Document) bool {
		return doc.Title == title
	})
	if !added {
		t.Errorf("новый документ не добавился в коллекцию документов")
	}
}

func Test_DeleteDoc(t *testing.T) {
	t.Run("Valid case", func(t *testing.T) {
		id := 444
		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/docs/%d", id), nil)
		rr := httptest.NewRecorder()

		api.Router.ServeHTTP(rr, req)
		if rr.Code != http.StatusAccepted {
			t.Errorf("код неверен: получили %d, хотели %d", rr.Code, http.StatusAccepted)
		}

		got := rr.Body.String()
		want := "Document successfully deleted!"
		if got != want {
			t.Errorf("\n получили: \n%s\n хотели: \n%s\n", got, want)
		}
		deleted := !slices.ContainsFunc(scanner.Docs, func(doc crawler.Document) bool {
			return doc.ID == id
		})
		if !deleted {
			t.Errorf("новый документ не удалился из коллекции")
		}
	})

	t.Run("ID not exist case", func(t *testing.T) {
		id := 666
		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/docs/%d", id), nil)
		rr := httptest.NewRecorder()

		api.Router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("код неверен: получили %d, хотели %d", rr.Code, http.StatusBadRequest)
		}

		got := rr.Body.String()
		want := fmt.Sprintf("Document with ID[%d] does not exist!", id)
		if got != want {
			t.Errorf("\n получили: \n%s\n хотели: \n%s\n", got, want)
		}
	})

	t.Run("Invalid id: int type overflow case", func(t *testing.T) {
		id := "9999999999999999999999999999999999"
		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/docs/%s", id), nil)
		rr := httptest.NewRecorder()

		api.Router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("код неверен: получили %d, хотели %d", rr.Code, http.StatusBadRequest)
		}

		got := rr.Body.String()
		want := fmt.Sprintf("ID[%s] is not a correct number", id)
		if got != want {
			t.Errorf("\n получили: \n%s\n хотели: \n%s\n", got, want)
		}
	})

}

func Test_UpdateDoc(t *testing.T) {
	id := 444
	title := "Updated"
	url := "http://updateddoc.com"
	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/docs/%d?title=%s&url=%s", id, title, url), nil)
	rr := httptest.NewRecorder()

	api.Router.ServeHTTP(rr, req)
	if rr.Code != http.StatusAccepted {
		t.Errorf("код неверен: получили %d, хотели %d", rr.Code, http.StatusAccepted)
	}

	got := rr.Body.String()
	want := "Document successfully updated!"
	if got != want {
		t.Errorf("\n получили: \n%s\n хотели: \n%s\n", got, want)
	}

	updated := slices.ContainsFunc(scanner.Docs, func(doc crawler.Document) bool {
		return doc.Title == title
	})
	if !updated {
		t.Errorf("данные о документе не поменялись внутри коллекции")
	}
}
