package main

import (
	"encoding/json"
	"fmt"
)

type Layer1 struct {
	Layer2
	Layer4 `json:"layer_4"`
}

type Layer2 struct {
	Layer3
}

type Layer3 struct {
	Name string `json:"name"`
}

type Layer4 struct {
	Age int `json:"age"`
}

func main() {
	l := Layer1{
		Layer2{
			Layer3{
				Name: "shou",
			},
		},
		Layer4{Age: 18},
	}
	data, _ := json.MarshalIndent(l, "", "  ")
	fmt.Printf("%v\n", string(data))

	var ll Layer1
	_ = json.Unmarshal(data, &ll)
	fmt.Printf("%+v\n", ll)
}
