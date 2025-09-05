package api

import (
	"encoding/json"
	"fmt"
	"go_search/13-homework/pkg/crawler"
	"go_search/13-homework/pkg/index"
	"go_search/13-homework/pkg/scanner"
	"math/rand/v2"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type API struct {
	Router *mux.Router
}

func New() *API {
	r := mux.NewRouter()
	api := &API{Router: r}

	r.HandleFunc("/docs", api.AddDoc).Methods("POST")
	r.HandleFunc("/docs/{id:[0-9]+}", api.DeleteDoc).Methods("DELETE")
	r.HandleFunc("/docs/{id:[0-9]+}", api.UpdateDoc).Methods("PATCH", "PUT", "POST")
	r.HandleFunc("/docs", api.ShowDocs).Methods("GET")
	r.HandleFunc("/search", api.Search).Methods("GET")

	return api
}

func (api *API) Search(w http.ResponseWriter, r *http.Request) {
	q := r.FormValue("q")
	if q == "" {
		w.Write([]byte("Empty searching query!"))
		return
	}
	res := index.Search(q)

	for _, id := range res {
		i, ok := slices.BinarySearchFunc(scanner.Docs, id, func(d crawler.Document, n int) int {
			return d.ID - n
		})
		if ok {
			fmt.Fprintf(w, `<a href="%s">%s</a><br>`, scanner.Docs[i].URL, scanner.Docs[i].Title)
		}
	}
}

func (api *API) ShowDocs(w http.ResponseWriter, r *http.Request) {
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

func (api *API) AddDoc(w http.ResponseWriter, r *http.Request) {
	title := strings.TrimLeft(r.FormValue("title"), " ")
	url := strings.TrimLeft(r.FormValue("url"), " ")
	if title == "" || url == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Required params are missing!\n(Title and URL should not be empty)"))
		return
	}
	doc := crawler.Document{ID: int(rand.Float32() * 10000), Title: title, URL: url}
	scanner.Mu.Lock()
	scanner.Docs = append(scanner.Docs, doc)
	slices.SortFunc(scanner.Docs, func(a crawler.Document, b crawler.Document) int {
		return a.ID - b.ID
	})
	scanner.Mu.Unlock()

	index.Mu.Lock()
	index.Add(doc)
	index.Mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Document successfully created!"))
}

func (api *API) DeleteDoc(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("ID[%s] is not a correct number", mux.Vars(r)["id"])))
		return
	}

	exists := slices.ContainsFunc(scanner.Docs, func(doc crawler.Document) bool {
		return doc.ID == id
	})

	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Document with ID[%d] does not exist!", id)))
		return
	}

	scanner.Mu.Lock()
	scanner.Docs = slices.DeleteFunc(scanner.Docs, func(doc crawler.Document) bool {
		return doc.ID == id
	})
	scanner.Mu.Unlock()

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Document successfully deleted!"))
}

func (api *API) UpdateDoc(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("ID[%s] is not a correct number", mux.Vars(r)["id"])))
		return
	}

	exists := slices.ContainsFunc(scanner.Docs, func(doc crawler.Document) bool {
		return doc.ID == id
	})

	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Document with ID[%d] does not exist!", id)))
		return
	}

	title := strings.TrimLeft(r.FormValue("title"), " ")
	url := strings.TrimLeft(r.FormValue("url"), " ")
	if title == "" || url == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Required params are missing!\n(Title and URL should not be empty)"))
		return
	}

	scanner.Mu.Lock()
	i := slices.IndexFunc(scanner.Docs, func(doc crawler.Document) bool {
		return doc.ID == id
	})
	doc := &scanner.Docs[i]
	doc.Title = title
	doc.URL = url
	scanner.Mu.Unlock()

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Document successfully updated!"))
}
