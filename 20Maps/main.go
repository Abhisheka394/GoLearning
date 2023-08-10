package main

import "fmt"

func main() {
	var myMap map[string]any
	myMap = make(map[string]any)
	myMap["522001"] = "G-1"
	myMap["522003"] = "G-2"
	myMap["574884"] = "D-2"
	myMap["572828"] = 485

	myMap["690034"] = 56.5443

	_, ok := myMap["522001"]
	if ok {
		fmt.Println("Key exists")
	} else {
		fmt.Println("Key exists")
	}

	delete(myMap, "522001")

	for key, val := range myMap {
		fmt.Println("Key: ", key, "value: ", val)
	}
}
