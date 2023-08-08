package main

import (
	"fmt"
	"reflect"
)

// Usually constraints are stored in data segments
// only at compile time value
const PI float32 = 3.14

const SQUARE_OF_PI1 = PI * PI //this is a valid const

const SQUARE_OF_PI2 = 3.14 * 3.14

const a = 18

const d = a * a

var pi float32 = 3.14

//const SQUARE_OF_PI3 = pi * pi

const (
	DAYS    = 7
	MONTHS  = 2
	HOURS   = 12
	MINUTES = 60
	SECONDS = 60
)

func main() {
	//PI=3.14
	fmt.Println(PI, SQUARE_OF_PI1, SQUARE_OF_PI2, reflect.TypeOf(a))

}
