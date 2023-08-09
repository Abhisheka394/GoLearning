package main

import "fmt"

// 5. Create Variable of all num types (all 14 data types + rune and byte), declare and assign values of all variables , perform addition and multiplicaton

func main() {
	var (
		num1 int   = 100
		num2 int8  = 125
		num3 int16 = 2345
		num4 int32 = 748480
		num5 int64 = 83833893993
		num6 uint8 = 93
		// num7   uint16  = 6750
		// num8   uint32  = 838339933
		num9   uint64  = 8383399399399939
		float1 float32 = 3993.3939
		float2 float64 = 883993333.399393

		// bool1 bool   = true
		// str1  string = "abcde"
		// byte1 byte   = 230
		// rune1 rune   = 2344555
	)

	var temp1 = num1 + int(num2)
	fmt.Println("temp1: ", temp1)

	var temp2 = num5 + int64(float1)
	fmt.Println("temp2: ", temp2)

	var temp3 = float64(num3) * float2 // converting int16 to float64
	fmt.Println("temp3: ", temp3)

	var temp4 = float32(num4) * float1
	fmt.Println("temp4: ", temp4)

	var temp5 = num4 * int32(float1) //converting float to int results in loss of value
	fmt.Println("temp5: ", temp5)

	var temp6 = uint64(num6) * num9
	fmt.Println("temp6: ", temp6)

}
