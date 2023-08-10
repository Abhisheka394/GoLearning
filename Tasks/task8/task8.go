package main

import (
	"fmt"
	"strings"
)

// Conditions
// State : Karnataka , AP, Delhi ,UP Gender: M,F Age: >0 Height:

// State	Gender	Height	Age	Ticket Status
// Karnataka	F			No ticket
// AP	F	<110cm	<5y	No ticket
// Delhi	F			No Ticket
// UP	F	<120cm	<6y	No ticket
// Karnataka	M	<110cm	<5y	No ticket
// AP	M	<110cm	<5y	No ticket
// Delhi	M	<130cm	<7y	No Ticket
// UP	M	<120cm	<6y	No ticket
// Other than the above table ,It is a full ticket

func main() {
	decideTicket("AP", "F", 140, 14)
	decideTicket("UP", "F", 108, 4)
	decideTicket("Karnataka", "F", 140, 14)
	decideTicket("Delhi", "F", 140, 14)
	decideTicket("Karnataka", "M", 140, 14)
	decideTicket("Delhi", "M", 90, 4)
	decideTicket("UP", "M", 140, 14)
	decideTicket("AP", "M", 110, 6)
}

func decideTicket(state string, gender string, height int, age int) {
	fmt.Println("State:", state, "gender:", gender, "height: ", height, "age: ", age)
	switch {
	case strings.ToLower(state) == "karnataka" && gender == "F":
		fmt.Println("No Ticket")
	case strings.ToLower(state) == "delhi" && gender == "F":
		fmt.Println("No Ticket")
	case strings.ToLower(state) == "ap" && gender == "F" && height < 110 && age < 5:
		fmt.Println("No Ticket")
	case strings.ToLower(state) == "up" && gender == "F" && height < 120 && age < 6:
		fmt.Println("No Ticket")
	case strings.ToLower(state) == "karnataka" && gender == "M" && height < 110 && age < 5:
		fmt.Println("No Ticket")
	case strings.ToLower(state) == "delhi" && gender == "M" && height < 110 && age < 7:
		fmt.Println("No Ticket")
	case strings.ToLower(state) == "ap" && gender == "M" && height < 130 && age < 5:
		fmt.Println("No Ticket")
	case strings.ToLower(state) == "up" && gender == "M" && height < 120 && age < 6:
		fmt.Println("No Ticket")
	default:
		fmt.Println("Full Ticket")
	}
}
