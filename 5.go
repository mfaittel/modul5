package main

import "fmt"

func ganjil(n, jadi int) {
	if jadi > n {
		return
	}

	fmt.Print(jadi, " ")

	ganjil(n, jadi+2)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	ganjil(n, 1)
	fmt.Println()
}
