package main

func main() {
outer:
	for i := 1; i == 1; {
		for j := 1; j <= 10; j++ {
			println("i:", i, "j:", j)
			if j == 10 {
				break outer
			}
		}
	}
	println("///////")
	for i, j := 0, 10; i <= j; i, j = i+2, j+1 {
		println("i:", i, "j:", j)

	}

}
