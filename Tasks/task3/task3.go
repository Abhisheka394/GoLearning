package main

import (
	"fmt"
	"math"
	"reflect"
)

// 3. Create constants and perform mathematical calculation on constants  Addition, Substraction, Division

func main() {
	const num1 = 5
	const num2 = 23
	const num3 = 4000
	const num4 = 4.566
	const num5 = 5.6766
	//addition
	const sum = num1 + num2
	fmt.Println("Sum: ", sum)

	// multiplicaton
	const prod = num1 * num2
	fmt.Println("Product: ", reflect.TypeOf(prod))

	const prod2 = num2 * num3 // <- This will give error as constant 800 of type int8) overflows int8
	fmt.Println(prod2)

	//const mod = math.Mod(num5, num4)
	fmt.Println("mod: ", math.Mod(num5, num4))
}
