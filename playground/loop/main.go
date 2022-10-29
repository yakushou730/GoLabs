package main

import "fmt"

func main() {
	ms := []int{100, 200, 300, 400}
	for i, v := range ms {
		fmt.Printf("i: %d, v: %d, &v: %X, &ms[k]: %X\n", i, v, &v, &ms[i])
	}
}
