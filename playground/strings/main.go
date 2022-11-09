package main

import (
	"fmt"
	"strings"
)

func main() {
	// trim
	fmt.Println(strings.TrimRight("12x3oxo", "xo"))
	fmt.Println(strings.TrimSuffix("123oxo", "xo"))

	// concatenation
	s := "Shou & Anna"
	sb := strings.Builder{}
	sb.Grow(len(s))
	sb.WriteString(s)
	fmt.Printf("%v\n", sb.String())
}
