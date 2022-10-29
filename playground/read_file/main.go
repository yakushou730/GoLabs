package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("test.txt")
	fmt.Printf("%v\n", string(data))
}
