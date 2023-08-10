package main

import "fmt"

// Variadic Parameter
// should be the last parameter
// there should be only one variadic parameter in a function

func main() {

	fmt.Println()
	arr1 := [...]int{1, 2, 3, 4, 5}
	slice1 := make([]int, 5)
	slice1 = arr1[:]
	var sum1 int = sumFunc(slice1...)
	fmt.Println("Sum of arr1: ", sum1)
	arr2 := [...]int{4, 4, 4, 5, 3, 4, 3}
	slice2 := arr2[:]
	sum2 := sumFunc(slice2...)
	fmt.Println("Sum of arr2: ", sum2)

}

func sumFunc(args ...int) int {
	var sum int = 0
	for _, val := range args {
		sum += val
	}
	return sum
}
