package main

import "fmt"

func main() {
	arr := [3][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}

	fmt.Println("Before Swapping matrix")
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Printf("%v ", arr[i][j])
		}
		println()
	}

	for i := 0; i < len(arr)/2; i++ {
		for j := 0; j < len(arr[0]); j++ {
			arr[i][j], arr[j][i] = arr[j][i], arr[i][j]
		}
	}

	fmt.Println("After transposition:")
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Printf("%v ", arr[i][j])
		}
		println()
	}
}
