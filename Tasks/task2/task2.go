package main

import "fmt"

//Use fmt packages like Println, Printf, Sprint, Sprintf, etc and identify the differences

func main() {
	var num1 int = 18

	fmt.Println(num1)      // fmt.Println formats using the default formats
	fmt.Printf("%v", num1) // fmt.Printf formats according to a format specifier

	var str1 string = fmt.Sprintf("%v", num1) //Sprintf formats according to a format specifier and returns the resulting string
	fmt.Printf("\nstr1: %v", str1)

	var str2 string = fmt.Sprint(num1)
	fmt.Printf("\nstr2: %v", str2) //Sprint formats using the default formats
}
