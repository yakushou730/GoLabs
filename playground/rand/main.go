package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		// 包含 0，不包含 5
		fmt.Println(rand.Intn(5))
	}
}
