package main

func main() {

	check(32)
	check(28)
	check(33)
	check(48)
	fallThroughWrongUse(34)
	fallThroughWrongUse(58)

}

func check(num int) {
	switch {
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, "is divisible by 2")
	default:
		println(num, "is not divisible by 2,4,8")
	}
}

func fallThroughWrongUse(num int) {
	println("Wrong Implementation of FallThrough")
	switch {
	case num%2 == 0:
		println(num, "is divisible by 2")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%8 == 0:
		println(num, "is divisible by 8")
	default:
		println(num, "is not divisible by 2,4,8")
	}

}
