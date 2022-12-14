package main

import "fmt"

func main() {
	sli := make([]int, 10)
	ma := make(map[string]int)

	fmt.Printf("sli: %v\n", sli)
	fmt.Printf("ma: %v\n", ma)

	addToSlice(sli, 1)
	addToMap(ma, "key", 1)

	fmt.Printf("sli: %v\n", sli)
	fmt.Printf("ma: %v\n", ma)

}

func addToSlice(s []int, item int) {
	// mutated
	s[0] = item

	// unmutated
	//s = append(s, item)
}

func addToMap(m map[string]int, s string, i int) {
	m[s] = i
}
