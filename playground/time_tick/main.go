package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.Tick(1 * time.Second)
	for {
		select {
		case <-ticker:
			fmt.Println("ping for every 1 seconds")
		}
	}
}
