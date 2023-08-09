package main

import "fmt"

// array type -> [n]type
// The type [n]T is an array of n values of type T

func main() {
	var arr1 [5]int
	fmt.Println(arr1)

	arr1[0] = 10
	arr1[1] = 11
	arr1[2] = 12
	arr1[3] = 13
	arr1[4] = 14

	var arr2 [5]bool
	fmt.Println(arr2)
	fmt.Println("Capacity of arr2: ", cap(arr2))

	// for i := 0; i < len(arr1); i++ {
	// 	fmt.Print(arr1[i], " ")
	// }

	for _, num := range arr1 {
		fmt.Print(num, " ")
	}
	var sum = sumOfArray(arr1)
	print("\nSum of arr1: ", sum)

	arr3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr3: ", arr3)
	arr4 := [...]int{4, 4, 4, 5, 3, 4, 3}
	fmt.Println("arr4: ", arr4)

	//arr5 := [2][]int{{1, 2}, {3, 4}}
	//fmt.Print(arr5)

	// for _, row := range arr5 {
	// 	for _, val := range row {
	// 		fmt.Println(val)
	// 	}
	// }

	arr6 := [2][][]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
	for _, arr3d := range arr6 {
		for _, row := range arr3d {
			for _, val := range row {
				fmt.Println(val)
			}
		}
	}
}

func sumOfArray(arr [5]int) int {
	var sum int = 0
	for _, val := range arr {
		sum += val
	}
	return sum
}
