package main

import "fmt"

// Create Company struct Id,Name, Website, Telehone, Fax, Address // Address must be a promoted filed

// Create Address struct City, Line1, Street, State, Country, PinCode , Map // Map must be a promoted field

// Create Map struct Lan, Lat

// -> Can you access Lan and Lap with the object of Company?

func main() {
	var m1 = Map{2384433.339399, 4488438.33838}
	addr1 := Address{"Blr", "23", "x street", "Karnataka", "India", "345678", m1}
	var c1 = Company{12, "Ab", "xyz.com", "2344555", "4455", addr1}
	fmt.Println(c1)

	var c2 = Company{12, "Ab", "abc.com", "5394945", "9647484", Address{"Chennai", "12", "church street", "TamilNadu", "India", "383993", Map{21292213.339399, 28273389.33838}}}
	fmt.Println(c2)

	c3 := Company{}
	c3.Id = 800
	c3.Name = "Roy"
	c3.Website = "r.com"
	c3.Telephone = "8098763"
	c3.Fax = "555434"
	c3.City = "Kochi"
	c3.Line1 = "34"
	c3.Street = "street"
	c3.State = "Kerala"
	c3.Country = "India"
	c3.Pincode = "4002992"
	c3.Lon = 48939.3839
	c3.Lat = 28282.383993
	fmt.Println(c3)
}

type Company struct {
	Id        int
	Name      string
	Website   string
	Telephone string
	Fax       string
	Address
}

type Address struct {
	City, Line1, Street, State, Country, Pincode string
	Map
}

type Map struct {
	Lon, Lat float32
}
