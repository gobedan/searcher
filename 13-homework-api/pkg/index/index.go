package index

import (
	"go_search/13-homework-api/pkg/crawler"
	"slices"
	"strings"
	"sync"
)

var Index = make(map[string][]int)
var Mu sync.Mutex

func Add(d crawler.Document) {
	words := strings.Split(d.Title, " ")

	for _, w := range words {
		wkey := strings.ToLower(w)
		if !slices.Contains(Index[wkey], d.ID) {
			Index[wkey] = append(Index[wkey], d.ID)
		}
	}
}

func Search(s string) []int {
	res := []int{}
	s = strings.ToLower(s)

	for w, id := range Index {
		if w == s {
			res = append(res, id...)
		}
	}
	return res
}
