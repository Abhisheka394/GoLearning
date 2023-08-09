package main

import (
	"fmt"
	"strings"
)

// 7. Conditional statements states - karnataka,AP, Delhi, UP, gender, age>0,  if state is karnataka or delhi and gender is female then no ticket, if state=k nd age>5 ,

func checkTicket(state string, gender string, height int, age int) {
	fmt.Println("Passenger Info-> State:", state, "Gender:", gender, "Height:", height, "Age:", age)
	if gender == "F" {
		if strings.ToLower(state) == "karanataka" || strings.ToLower(state) == "delhi" {
			fmt.Println("No Ticket")
		} else if strings.ToLower(state) == "ap" && height < 110 && age < 5 {
			fmt.Println("No Ticket")
		} else if strings.ToLower(state) == "up" && height < 120 && age < 6 {
			fmt.Println("No Ticket")
		} else {
			fmt.Println("Full Ticket")
		}
	} else if gender == "M" {
		if strings.ToLower(state) == "karnataka" && height < 110 && age < 5 {
			fmt.Println("No Ticket")
		} else if strings.ToLower(state) == "ap" && height < 110 && age < 5 {
			fmt.Println("No Ticket")
		} else if strings.ToLower(state) == "delhi" && height < 130 && age < 7 {
			fmt.Println("No Ticket")
		} else if strings.ToLower(state) == "up" && height < 120 && age < 6 {
			fmt.Println("No Ticket")
		} else {
			fmt.Println("Full Ticket")
		}
	} else {
		fmt.Println("Invalid info,  so Full Ticket")
	}
}

func main() {
	checkTicket("AP", "F", 150, 24)
	checkTicket("Delhi", "M", 90, 4)
	checkTicket("Karnataka", "F", 150, 24)
	checkTicket("Karnataka", "M", 130, 8)
	checkTicket("UP", "F", 150, 24)
	checkTicket("delhi", "F", 120, 9)

}
