package main

import (
	"fmt"
	"reflect"
)

func main() {
	r1 := Rect{L: 2, B: 3.54}
	fmt.Println(AreaOfRect(r1))
	fmt.Println(PeriOfRect(r1))

	r2 := new(Rect)
	r2.L = 49.330
	fmt.Println(reflect.TypeOf(r2))

}

type Rect struct {
	L, B float32
	A, P float32
}

type Square struct {
	S float32
}

// To write Methods , we need to add a receiver

func (r Rect) PeriBy() float32 {
	r.P = 2 * (r.B + r.L)
	return r.P

}

func (r Rect) AreaBy() float32 {
	r.A = r.B * r.L
	return r.A

}

func (r *Rect) Peri() float32 {
	r.P = 2 * (r.B + r.L)
	return r.P

}

func (r *Rect) Area() float32 {
	r.A = r.B * r.L
	return r.A

}

func AreaOfRect(r Rect) float32 { //function
	return r.L * r.B
}

func PeriOfRect(r Rect) float32 {
	return 2 * (r.L + r.B)
}
