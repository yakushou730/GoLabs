package main

import (
	"fmt"
	"time"
)

func main() {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("reach 3 sec...")
	}
}
