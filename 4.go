package main

import "fmt"

func barisirab(n, jadi int) {
	fmt.Print(jadi, " ")

	if jadi > 1 {
		barisirab(n, jadi-1)
		fmt.Print(jadi, " ")
	}
}

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	barisirab(n, n)
	fmt.Println()
}
