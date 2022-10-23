package main

import "fmt"

type FunctionAdd1 func(int, int) int

func Add(i, j int) int {
	return i + j
}

func (f FunctionAdd1) AddExtra1(i, j int) int {
	return f(i, j) + 1
}

func main() {
	ff := FunctionAdd1(Add)
	fmt.Println(ff.AddExtra1(1, 2))
}
