package main

import "fmt"

func main() {
	fmt.Println(reverseString("hello"))
	fmt.Println(reverseString("Shou"))
	fmt.Println(reverseString("中文測試啦"))
}

func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
