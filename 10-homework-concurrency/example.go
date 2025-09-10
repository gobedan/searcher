package main

import (
	"fmt"
	"sync"
)

func worker(n int) chan int {
	ch := make(chan int)
	// эта горутина для канала из return
	go func() {
		// выдаем задания - закрываем, когда его прочитали
		defer close(ch)
		ch <- n * n
	}()
	return ch
}

func fanIn[T any](channels ...<-chan T) <-chan T {
	ch := make(chan T)
	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, c := range channels {
		// эти горутины для каналов из аргументов
		go func(in <-chan T) {
			// синхронизируемся только когда входящий канал закрыт (мы считали все задания) + исходящий канал прочитан (вызывающий обработчик все обработал)
			defer wg.Done()
			for i := range in {
				ch <- i
			}
		}(c)
	}

	// ждет в отдельной горутине - иначе дедлок - нужно вернуть канал туда, где его сперва будут читать
	// эта горутина для канала, который мы возвращаем в return
	go func() {
		// ждет закрытия ВСЕХ входящих каналов (мы считали все задания всех источников) и окончания чтения из исходящего канала (вызывающий код обработал все результаты всех источников)
		wg.Wait()
		close(ch)
	}()

	return ch
}

func main() {
	var chans []<-chan int
	for i := 0; i < 10; i++ {
		chans = append(chans, worker(i))
	}
	// горутины для chans.. и ch уже есть внутри функции fanIn
	ch := fanIn(chans...)
	for val := range ch {
		fmt.Println(val)
	}
}
