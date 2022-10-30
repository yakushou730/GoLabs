package main

import "fmt"

func main() {
	ss := []string{}
	var sss []string
	fmt.Printf("%v\n", ss)
	fmt.Printf("%p\n", &ss)
	fmt.Println("--------")
	fmt.Printf("%v\n", sss)
	fmt.Printf("%p\n", &sss)
}
