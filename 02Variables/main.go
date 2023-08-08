package main

import (
	"fmt"
	"reflect"
)

func main() {
	var age uint8 = 38
	fmt.Printf("Age: %v && Type of age: %s", age, reflect.TypeOf(age))

	num1 := 123
	fmt.Println("\n", num1)

	var num6 int
	var num7 float32
	var num8 bool
	var num9 string
	fmt.Printf("%v %v %v %v", num6, num7, num8, num9)
}
