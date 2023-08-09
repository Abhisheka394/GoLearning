package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	var sum, err = add(2, 4555)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sum:", sum)
	}
}

func add(a, b any) (float32, error) {
	if a == nil || b == nil {
		return -1, errors.New("Either a or b is null")
	}
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return -1, errors.New("a and b are of Different types")
	}
	var sum float32
	switch a.(type) {
	case int:
		sum = float32(a.(int)) + float32(b.(int))
	case int8:
		sum = float32(a.(int8)) + float32(b.(int8))
	case int16:
		sum = float32(a.(int16)) + float32(b.(int16))
	case int32:
		sum = float32(a.(int32)) + float32(b.(int32))
	case int64:
		sum = float32(a.(int64)) + float32(b.(int64))
	case uint8:
		sum = float32(a.(uint8)) + float32(b.(uint8))
	case uint16:
		sum = float32(a.(uint16)) + float32(b.(uint16))
	case uint32:
		sum = float32(a.(uint32)) + float32(b.(uint32))
	case uint64:
		sum = float32(a.(uint64)) + float32(b.(uint64))
	case float32:
		sum = a.(float32) + a.(float32)
	case float64:
		sum = float32(a.(float64)) + float32(a.(float64))
	default:
		return -1, errors.New("Error")
	}
	return sum, nil
}

// func addAny(a, b any) (float32, error) {
// 	if reflect.TypeOf(a) != reflect.TypeOf(b) {
// 		return -1, errors.New("a and b are of Different types")
// 	}
// 	var sum float32
// 	var a1 = reflect.TypeOf(a)
// 	var a2 = a1.Kind()
// 	var b1 = reflect.TypeOf(b)
// 	var b2 = b1.Kind()
// 	sum = float32(a.(a2)) + float32(b.(b2))
// 	return sum, nil
// }
