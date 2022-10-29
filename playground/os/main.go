package main

import (
	"fmt"
	"os"
)

func main() {
	_ = os.Setenv("NAME", "gopher")
	_ = os.Setenv("BURROW", "/usr/gopher")
	fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}."))
}
