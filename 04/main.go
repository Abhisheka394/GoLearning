package main

import "fmt"

type integer = int

func main() {

	var num1 integer = 10390303

	var num2 rune = 'A'

	var num3 int32 = 'B'

	var num4 uint8 = '#'
	var num5 byte = 'B'

	fmt.Printf("num1:%v, num2:%v, num3:%v, num4:%v, num5:%v", num1, num2, num3, num4, num5)
}
