package main

import (
	"flag"
	"fmt"
	"go_search/02-homework-syntax/pkg/crawler"
	"go_search/02-homework-syntax/pkg/crawler/spider"
	"strings"
)

/*
	Задание

# Задача №2

Создайте исполняемый пакет, который бы при запуске использовал пакет «crawler» для сканирования сайтов «go.dev» и «golang.org»

Результат сканирования сайтов нужно объединить.

# Задача №3

После окончания сканирования, если пользователь ввел какое-либо слово в флаге «s», приложение должно напечатать все ссылки, где это слово встречается.

Для обработки флагов вызова используйте пакет «flag»

Пример вызова программы:
gosearch -s documents
*/
const depth = 2

var urls = []string{"https://go.dev", "https://golang.org"}

var sFlag = flag.String("s", "", "keywords to search for IN TITTLE")

var docs = []crawler.Document{}

func main() {
	flag.Parse()
	*sFlag = strings.TrimSpace(*sFlag)

	for _, u := range urls {
		search(u)
	}

	for _, doc := range docs {
		fmt.Printf("%+v\n", doc)
	}

	if *sFlag != "" {
		fmt.Printf("Searching matchers for phrase: %s\n", *sFlag)
		count := 0
		for _, doc := range docs {
			if strings.Contains(strings.ToLower(doc.Title), strings.ToLower(*sFlag)) {
				fmt.Printf("Match found in: %s\t-\t%s\n", doc.URL, doc.Title)
				count++
			}
		}
		fmt.Printf("Total matches: %d\n", count)
	}
}

func search(url string) {
	spd := spider.New()
	scans, err := spd.Scan(url, depth)
	if err != nil {
		fmt.Printf("Scan failed for url:%s\n\tError:%v", url, err)
		return
	}
	docs = append(docs, scans...)
}
