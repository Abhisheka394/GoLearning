package main

import (
	"errors"
	"fmt"
)

func main() {
	mp := map[int]string{

		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}
	fmt.Println(mp)
	println()

	val, err := Delete(mp, 90)
	if err != nil {
		fmt.Println(err)
		fmt.Println(mp)
	} else if val {
		fmt.Println("Record deleted succesfully!!")
		fmt.Println(mp)
	}
	println()

	val, err = Update(mp, 91, "Horses")
	if err != nil {
		fmt.Println(err)
		fmt.Println(mp)
	} else if val {
		fmt.Println("Record updatd succesfully!!")
		fmt.Println(mp)
	}

}

func Delete(mp map[int]string, key int) (bool, error) {
	_, ok := mp[key]
	if mp == nil {
		return false, errors.New("Map is Nil")
	} else if ok {
		delete(mp, key)
		return true, nil
	} else {
		return false, nil
	}
}

func Update(mp map[int]string, key int, val string) (bool, error) {
	_, ok := mp[key]
	if mp == nil {
		return false, errors.New("Map is Nil")
	} else if !ok {
		return false, errors.New("Key does no exist")
	} else {
		mp[key] = val
		return true, nil
	}
}
