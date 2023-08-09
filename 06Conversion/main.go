package main

import "fmt"

func main() {
	var num1 int = 65001

	var num2 uint8 = uint8(num1) //

	fmt.Println("num1: ", num2)

	var num3 uint8 = 43
	var num4 int = int(num3)
	fmt.Println("num3: ", num4) //

	var str1 string = "H"

	var num5 int = int(str1[0])
	fmt.Println(num5)

	fmt.Println(int('A'))

	var num6 int = 66
	str3 := fmt.Sprint(num6)
	fmt.Printf(str3)

}
