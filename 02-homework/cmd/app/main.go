package main

import (
	"flag"
	"fmt"
	"go_search/02-homework/pkg/crawler"
	"go_search/02-homework/pkg/crawler/spider"
	"strings"
)

/*
Задание №1

Создать обратный индекс для документов.

Разместить индекс в пакете «index».

Ключом индекса будет каждое слово из описания ссылок, значением – номер документа.

После сканирования сайтов все найденные ссылки должны храниться в виде структуры следующего вида:
// Document - документ, веб-страница, полученная поисковым роботом.
type Document struct {
ID int
URL string
Title string
}

Проиндексированные документы должны храниться в массиве документов и объекте из пакета «index».
При этом для каждого документа нужно добавить его номер – поле в структуре данных вместе с URL и Title.
Массив документов должен быть отсортирован по номерам, используя сортировку из стандартной библиотеки.

# Задача №2

Переделать метод поисковой выдачи на использование индекса.

Для поиска по индексу приложение должно импортировать поисковый индекс из пакета «index» как зависимость и искать по индексу,
а не по массиву документов.

Поиск по индексу должен выдавать номера документов (поле в структуре данных документа).

# Задача №3

После получения номеров документов из индекса нужно использовать бинарный поиск по отсортированному ранее массиву документов
(по номерам документов).

Можно воспользоваться стандартной библиотекой или написать свою реализацию.
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
