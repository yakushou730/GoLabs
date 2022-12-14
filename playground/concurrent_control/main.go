package main

import (
	"fmt"
	"sync"
	"time"
)

var token = make(chan int, 3)

func main() {
	jobCount := 30
	var wg sync.WaitGroup

	for i := 0; i < jobCount; i++ {
		wg.Add(1)
		go runJob(i, &wg)
	}

	wg.Wait()
	fmt.Println("Finish all jobs !!")
}

func runJob(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	token <- 1
	fmt.Printf("running job id %d\n", id)
	time.Sleep(time.Second)
	<-token
}
