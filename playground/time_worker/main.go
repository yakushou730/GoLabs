package main

import (
	"fmt"
	"time"
)

func main() {
	taskCh := make(chan int, 100)

	go worker(taskCh)

	for i := 0; i < 30; i++ {
		taskCh <- i
	}

	select {
	case <-time.After(time.Hour):
	}
}

func worker(tashCh <-chan int) {
	const N = 5

	for i := 0; i < N; i++ {
		go func(id int) {
			for {
				task := <-tashCh
				fmt.Printf("finish task: %d by worker %d\n", task, id)
				time.Sleep(time.Second)
			}
		}(i)
	}
}
