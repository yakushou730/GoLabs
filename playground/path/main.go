package main

import (
	"fmt"
	"path"
)

func main() {
	res := path.Ext("/foo/bar/test.css")
	fmt.Println(res)
}
