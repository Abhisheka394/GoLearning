package main

import "fmt"

func main() {
	var p1 Person
	p1.Id = 100
	p1.Name = "Victoria"
	p1.Email = "Victoria@VS.com"
	p1.Mobile = "93833993"
	addr1 := Address{"23", "Kochi", "Church Street", "Kerala", "India", "688599"}
	p1.Addr = addr1
	fmt.Println(p1)
	addr2 := Address{"23", "Blr", "Church Street", "Karnataka", "India", "889233"}
	p2 := Person{109, "VS", "VS@s.com", "8383939", addr2}
	fmt.Println(p2)

	p3 := Person{}
	p3.Id = 133
	p3.Name = "Secret"
	p3.Email = "Secret@Secret.com"
	p3.Mobile = "3838939"
	p3.Addr = addr2
	p3.Addr.Line1 = "345"
	fmt.Println(p3)
}

type Person struct {
	Id     int
	Name   string
	Email  string
	Mobile string
	Addr   Address
}

type Address struct {
	Line1, City, Street, State, Country, PinCode string
}
