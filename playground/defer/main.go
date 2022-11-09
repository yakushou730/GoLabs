package main

import "fmt"

func main() {
	i := 0
	j := 0
	defer func(i int) {
		fmt.Println(i, j)
	}(i) // Passes i to the closure (evaluated right away)
	i++
	j++

	// 呼叫 defer 回馬上固定輸入值
	// 解決方案: 輸入指標 / 或用 closure
}
