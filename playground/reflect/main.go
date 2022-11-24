package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 5
	Print(a)

	b := &a
	Print(b)

	c := []string{"test"}
	Print(c)

	d := map[string]string{"a": "b"}
	Print(d)

	//
	x := 5
	y := &x
	v := reflect.ValueOf(y).Elem()
	fmt.Printf("%v %T\n", v.Int(), v.Int())
	v.SetInt(100)
	fmt.Printf("%v %T\n", v.Int(), v.Int())
	fmt.Printf("%v %T\n", v.Interface(), v.Interface())
	fmt.Printf("%v %T\n", *y, *y)

	//

	u1 := User{
		Name: "Shou",
		Age:  20,
	}
	PrintStruct(u1)

	vv := reflect.ValueOf(&u1).Elem()
	vv.FieldByName("Name").SetString("Anna")
	vv.FieldByName("Age").SetInt(15)
	PrintStruct(u1)
}

func Print(i any) {
	fmt.Println("Type:", reflect.TypeOf(i))
	fmt.Println("Value:", reflect.ValueOf(i))
}

type User struct {
	Name string `des:"userName"`
	Age  int    `des:"userAge"`
}

func PrintStruct(s any) {
	sT := reflect.TypeOf(s)
	sV := reflect.ValueOf(s)

	fmt.Printf("type %s %v {\n", sT.Name(), sT.Kind().String())

	for i := 0; i < sT.NumField(); i++ {
		field := sT.Field(i)
		value := sV.Field(i)

		fmt.Printf("\t%s\t%v\t= %v\t (description: %s)\n", field.Name, field.Type.String(), value.Interface(), field.Tag.Get("des"))
	}

	fmt.Println("}")
}
