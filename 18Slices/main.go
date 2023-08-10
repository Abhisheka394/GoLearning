package main

import "fmt"

func main() {

	slice1 := []int{10, 20, 30, 40, 50, 60}
	fmt.Println("Slice1: ", slice1)
	arr1 := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("arr1: ", arr1)
	slice2 := arr1[:5]
	fmt.Println("Slice2: ", slice2)
	slice3 := arr1[2:]
	fmt.Println("Slice3: ", slice3)

	slice4 := make([]int, 5, 10)
	slice4[3] = 5
	slice4[4] = 80
	fmt.Println("slice4: ", slice4)

	slice5 := slice4
	slice4 = append(slice4, 6)
	fmt.Println("slice4: ", slice4, "  slice5: ", slice5)
	slice5 = append(slice5, 16)
	fmt.Println("slice4: ", slice4, "  slice5: ", slice5)
	slice5 = append(slice5, 7)
	fmt.Println("slice4: ", slice4, "  slice5: ", slice5)
	slice4[0] = 89
	slice4[2] = 100
	fmt.Println("slice4: ", slice4)

	fmt.Println("slice5: ", slice5)
	fmt.Println(cap(slice4), cap(slice5))

	var slice6 = make([]int, len(slice5))
	slice6 = slice5
	fmt.Println(slice6)

}
