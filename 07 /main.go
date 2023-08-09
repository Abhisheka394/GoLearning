package main

import (
	"fmt"
	"reflect"
)

func main() {

	var iface1 any
	fmt.Println("iface1: ", iface1, "type: ", reflect.TypeOf(iface1)) // empty interface value is nil, as it is not pointing to any memory

	iface1 = 1000
	fmt.Println("iface1: ", iface1, "type: ", reflect.TypeOf(iface1))

	var int1 int
	int1 = iface1.(int)
	fmt.Println("int1: ", int1)

	var float1 float32 = 10.1234
	iface1 = float1
	fmt.Println("iface1: ", iface1, "type: ", reflect.TypeOf(iface1))

	var bool1 bool = true
	iface1 = bool1
	fmt.Println("iface1: ", iface1, "type: ", reflect.TypeOf(iface1))

	var bool2 bool
	bool2 = iface1.(bool)
	fmt.Println("bool2: ", bool2)

	var str1 string = "abcd"
	iface1 = str1
	fmt.Println("iface1: ", iface1, "type: ", reflect.TypeOf(iface1))

	var str2 string
	str2 = iface1.(string)
	fmt.Println("str2: ", str2)
}
