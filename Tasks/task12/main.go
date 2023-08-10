package main

import "fmt"

func main() {
	arr1 := []int{1, 32, 33, 234, 5, 2, 44, 22, 4, 3}
	fmt.Println("Second biggest element: ", secondBig(arr1))
	fmt.Println("Second Smallest element: ", secondSmall(arr1))

}

func secondBig(arr []int) int {
	var big1, big2 = arr[0], arr[0]
	for _, val := range arr {
		if val > big1 {
			big2 = big1
			big1 = val
		} else if val > big2 {
			big2 = val
		}
	}
	return big2
}

func secondSmall(arr []int) int {
	var small1, small2 = 100000, 100000
	for _, val := range arr {
		if val < small1 {
			small2 = small1
			small1 = val
		} else if val < small2 {
			small2 = val
		}
	}
	return small2
}
