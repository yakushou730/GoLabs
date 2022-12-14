package main

import "fmt"

type printer interface {
	print()
}

type user struct {
	name string
}

func (u user) print() {
	fmt.Printf("User Name: %s\n", u.name)
}

func main() {
	u := user{name: "Shou"}

	// 第一個塞進去的是 u 的 copy，所以即使後來改了資料也不會變
	entities := []printer{
		u,
		&u,
	}

	u.name = "Anna"

	for _, e := range entities {
		e.print()
	}
}
