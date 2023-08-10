package main

func main() {

}

type Comapny struct {
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
}
