package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	wg.Add(1)
	go publish(ctx, &wg)

	wg.Wait()

	fmt.Println("finished")
}

func publish(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("timeout: %v\n", ctx.Err())
			return
		default:
			fmt.Println("continue...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
