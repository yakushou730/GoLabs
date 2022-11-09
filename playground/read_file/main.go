package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("test.txt")
	fmt.Printf("%v\n", string(data))

	//
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//scanner := bufio.NewScanner(file)
}
