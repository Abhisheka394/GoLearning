package main

func main() {
	i := 1
loop:
	if i%2 == 0 {
		goto printme
	}
back:
	i++
	if i < 10 {
		goto loop
	}

printme:
	println(i, "is divissible by 2")
	if i < 10 {
		goto back
	}

}
