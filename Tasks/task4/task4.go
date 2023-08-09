package main

import "fmt"

var num1 int = 4

func temp() {
	//global variables can be accessed from anywhere
	fmt.Println("num1: ", num1)
}

func main() {
	temp()
	//global variables can be accessed from anywhere
	num1 = 56
	fmt.Println("num1: ", num1)
	var num2 = 48
	fmt.Println("num2: ", num2)
}
