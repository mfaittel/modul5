package main

import "fmt"

func faktor(n int, jadi int) {
	if jadi > n {
		return
	}
	if n%jadi == 0 {
		fmt.Print(jadi, " ")
	}
	faktor(n, jadi+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	fmt.Print("Faktor dari", n, "adalah: ")
	faktor(n, 1)
	fmt.Println()
}
