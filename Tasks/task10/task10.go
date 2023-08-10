package main

import "fmt"

func main() {
	arr := [3][]int{{2, 3, 4}, {3, 4, 5}, {45, 67, 9}}
	calRowSum(arr)
	calColSum(arr)
}

func calRowSum(arr [3][]int) {
	for i, row := range arr {
		var sum int = 0
		for _, val := range row {
			sum += val
		}
		fmt.Println("Row: ", i+1, " Sum: ", sum)
	}
}

func calColSum(arr [3][]int) {
	for i := 0; i < len(arr); i++ {
		var sum int = 0
		for j := 0; j < len(arr[0]); j++ {
			sum += arr[j][i]
		}
		fmt.Println("Col: ", i+1, " Sum: ", sum)
	}
}
