package main

import "fmt"

// In a given array find the biggest and smallest number in an array.
// Write a function getBigAndSmall(arr [10]int)(int,int)

func main() {
	arr1 := [...]int{1, 3, 2, 45, 6, 76, 87, 4, 323, 44}
	var big1, small1 = getBigAndSmall(arr1)
	fmt.Println("arr1: ", arr1)
	fmt.Println("Big1: ", big1, "&& Small1: ", small1)

	arr2 := [...]int{10, -3, 2, -45, 6, 76, 7, 4, 23, 44}
	var big2, small2 = getBigAndSmall(arr2)
	fmt.Println("arr1: ", arr2)
	fmt.Println("Big1: ", big2, "&& Small1: ", small2)
}

func getBigAndSmall(arr [10]int) (int, int) {
	var big, small = arr[0], arr[0]
	for _, val := range arr {
		if val < small {
			small = val
		}
		if val > big {
			big = val
		}
	}
	return big, small
}
