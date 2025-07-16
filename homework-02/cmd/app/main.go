package main

import (
	"flag"
	"fmt"
	"go_search/homework-02/pkg/crawler"
	"go_search/homework-02/pkg/crawler/spider"
	"log"
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
var sFlag = flag.String("s", "", "keywords to search for IN TITTLE")

var docs = []crawler.Document{}

func init() {
	flag.Parse()
	*sFlag = strings.TrimSpace(*sFlag)
}

func main() {
	search("https://go.dev")
	search("https://golang.org")
	for _, doc := range docs {
		fmt.Printf("%+v\n", doc)
	}

	if *sFlag != "" {
		fmt.Printf("Searching matchers for phrase: %s\n", *sFlag)
		count := 0
		for _, doc := range docs {
			if strings.Contains(doc.Title, *sFlag) {
				fmt.Printf("Match found in: %s\t-\t%s\n", doc.URL, doc.Title)
				count++
			}
		}
		fmt.Printf("Total matches: %d\n", count)
	}
}

func search(url string) {
	spd := spider.New()
	scans, err := spd.Scan(url, 2)
	if err != nil {
		log.Fatalf("Scan failed for url:%s\n\tError:%v", url, err)
	}
	docs = append(docs, scans...)
}
