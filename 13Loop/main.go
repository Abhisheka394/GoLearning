package main

import (
	"fmt"
)

func main() {
	println("Numbers from 1 to 10")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v ", i)
	}
	println("\nEven Numbers:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%v  ", i)
	}
	println()
	checkPrime(12)
	checkPrime(7)
	checkPrime(89)
}

func checkPrime(num int) {
	var isPrime = true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			fmt.Println(num, " is not a Prime Number")
			break
		}
	}
	if isPrime {
		fmt.Println(num, " is a Prime Number")
	}
}
