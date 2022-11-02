package main

import (
	"fmt"
	"sync"
)

// go run -race main.go

var msg string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()

	m.Lock()
	msg = s
	m.Unlock()
}

func main() {
	msg = "Hello World"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Hello Shou", &mutex)
	go updateMessage("Hello Anna", &mutex)

	wg.Wait()

	fmt.Println(msg)
}
