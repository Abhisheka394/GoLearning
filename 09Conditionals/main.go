package main

import "fmt"

func main() {
	var num1 = 12

	if num1%2 == 0 {
		fmt.Printf("%v is divisible by 2!", num1)
	} else {
		fmt.Printf("%v is not divisible by 2!", num1)
	}

	if num1%2 == 0 && num1%3 == 0 && num1%6 == 0 {
		fmt.Printf("\n%v is divisible by 2,3 and 6!", num1)
	} else {
		fmt.Printf("\n%v is not divisible by 2,3 and 6!", num1)
	}

	value := 55

	if value < 0 {
		fmt.Println("Less than 0")
	} else if value < 50 {
		fmt.Println("Grade C")
	} else if value < 100 {
		fmt.Println("Grade B")
	} else {
		fmt.Println("Grade A")
	}
}
