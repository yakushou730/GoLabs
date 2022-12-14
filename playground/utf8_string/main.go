package main

import "fmt"

func main() {
	s := "hello 世界"

	fmt.Println("for range...")
	for i, v := range s {
		fmt.Printf("%d: %v\n", i, string(v))
	}

	fmt.Println()

	fmt.Println("for i...")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d: %v\n", i, string(s[i]))
	}
}
