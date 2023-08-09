package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		num1   int     = 100
		num2   int8    = 125
		num3   int16   = 2345
		num4   int32   = 748480
		num5   int64   = 83833893993
		num6   uint8   = 93
		num7   uint16  = 6750
		num8   uint32  = 838339933
		num9   uint64  = 8383399399399939
		float1 float32 = 3993.3939
		float2 float64 = 883993333.399393

		bool1 bool   = true
		str1  string = "abcde"
		byte1 byte   = 230
		rune1 rune   = 2344555
	)

	var iface1, iface2, iface3, iface4 any

	iface1 = num1
	fmt.Println("iface1: ", iface1, "type: ", reflect.TypeOf(iface1))

	iface2 = float1
	fmt.Println("iface1: ", iface1, "type: ", reflect.TypeOf(iface1))

	var sum int = iface1.(int) + int(num3) + int(num6) + int(num7) + int(num8) + int(num9) + int(float1) + int(float2) // type asserting the interface
	fmt.Println("Num: ", sum)

	var prod float64 = iface2.(float64) + float64(num2) + float64(num4) + float64(num5) + float64(float2) + float64(byte1) + float64(rune1)
	fmt.Println("Prod: ", prod)

	iface3 = bool1
	var bool2 bool = iface3.(bool)
	fmt.Println("bool2: ", bool2)

	iface4 = str1
	var str2 string = iface4.(string)
	fmt.Println("str2: ", str2)

}
