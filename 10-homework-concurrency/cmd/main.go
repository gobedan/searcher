package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const MAX_SCORE = 5

func main() {
	ball := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(2)
	fmt.Println("==START==")
	go play(ball, "player1", wg)
	go play(ball, "player2", wg)
	ball <- "begin"
	wg.Wait()
	fmt.Println("==FINISH==")
}

func play(ch chan string, player string, wg *sync.WaitGroup) {
	var score int
	for hit := range ch {
		switch hit {
		case "begin":
			ch <- "ping"
			fmt.Printf("%s: ping\n", player)
		case "ping":
			if rand.Int()%100 > 20 {
				ch <- "pong"
				fmt.Printf("%s: pong\n", player)
			} else {
				score++
				if score < MAX_SCORE {
					ch <- "stop"
					fmt.Printf("%s hit! score: %d\n", player, score)
				} else {
					fmt.Printf("%s WON!\n", player)
					close(ch)
				}
			}
		case "pong":
			if rand.Int()%100 > 20 {
				ch <- "ping"
				fmt.Printf("%s: ping\n", player)
			} else {
				score++
				if score < MAX_SCORE {
					ch <- "stop"
					fmt.Printf("%s hit! score: %d\n", player, score)
				} else {
					fmt.Printf("%s WON!\n", player)
					close(ch)
				}
			}
		case "stop":
			ch <- "begin"
			fmt.Printf("%s: begin\n", player)
		}
	}
	fmt.Printf("%s finished with score: %d\n", player, score)
	wg.Done()
}
