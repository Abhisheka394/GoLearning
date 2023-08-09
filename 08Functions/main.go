package main

import "fmt"

func greet() {
	fmt.Println("Hello World!!")
}

func greetPerson(name string) {
	fmt.Println("Hello", name)
}

func greetAndMsg(name string, msg string) {
	fmt.Println("Hello", name, "!", msg)
}

func areaRect(l float32, b float32) float32 {
	return l * b
}

func areaAndPerimeter(l float32, b float32) (area float32, peri float32) {
	return l * b, 2 * (l + b)
}

func increment(num *int) {
	*num = *num + 1
}

func main() {
	greet()
	greetPerson("ADAM")
	greetAndMsg("Victoria", "How was ur day?")

	var areaRactange = areaRect(90.5, 78.567)
	fmt.Println("Area of rectangle:", areaRactange)

	var area, peri = areaAndPerimeter(3.4, 6.5)
	fmt.Println("Area:", area, "Perimeter:", peri)

	var num1 = 100
	increment(&num1)
	fmt.Println("num1:", num1)
}
