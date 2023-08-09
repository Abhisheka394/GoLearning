package main

func main() {
	var weekDay int = 5
	switch weekDay {
	case 1:
		println("Sunday")
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Saturday")
	default:
		println("NA")
	}

	var num int = 112

	switch {
	case num > 0 && num < 50:
		println("num is between 0 ad 50")
	case num >= 50 && num < 100:
		println("num is between 50 and 100")
	case num >= 100 && num < 200:
		println("num is between 100 and 200")
	default:
		println("Num is 200 or more")
	}

	//
}
