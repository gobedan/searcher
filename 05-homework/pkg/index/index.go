package index

import (
	"go_search/05-homework/pkg/crawler"
	"slices"
	"strings"
)

var Index = make(map[string][]int)

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

	for w, id := range Index {
		if w == s {
			res = append(res, id...)
		}
	}
	return res
}
