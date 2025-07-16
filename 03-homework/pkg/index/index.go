package index

import (
	"math/rand"
	"slices"
	"strings"
)

type Document struct {
	ID    int
	URL   string
	Title string
}

var Docs = []Document{}

var Index = make(map[string]int)

// ? так и должно затирать в индексе номера документов для повторяющихся слов?
func Add(d Document) {
	d.ID = int(rand.Float32() * 10000)
	Docs = append(Docs, d)

	words := strings.Split(d.Title, " ")
	for _, w := range words {
		Index[strings.ToLower(w)] = d.ID
	}
	slices.SortFunc(Docs, func(a Document, b Document) int {
		return a.ID - b.ID
	})
}
