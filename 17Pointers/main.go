package main

import (
	"errors"
	"fmt"
)

func main() {
	var num1 int = 100
	var ptr1 *int = &num1

	*ptr1 = 200
	var err = increment(&num1) // increment(ptr1) increment(*ptr2)  also works
	if err == nil {
		fmt.Println(err)
	}
	fmt.Println("value of num1: ", num1, "address 0f num1: ", &num1)
	fmt.Println("Value of num1 using ptr1: ", *ptr1, "address of num1: ", ptr1)

	var ptr2 **int = &ptr1
	fmt.Println("Value of num1 using ptr2: ", *ptr2, "address of num1: ", ptr2)

	var ptr3 *int
	if ptr3 == nil {
		fmt.Println("ptr3 is  a nil ptr")
	}

	var ptr4 = new(int)
	fmt.Println(*ptr4)
	*ptr4 = 500
	fmt.Println(*ptr4)

	var ptr5 = new(any)
	fmt.Println(*ptr5)
	*ptr5 = 5899.444
	fmt.Println(*ptr5)
	*ptr5 = "fjuufur"
	fmt.Println(*ptr5)

}

func increment(num *int) error {
	if num == nil {
		return errors.New("Nil Pointer")
	}
	*num++
	return nil
}
