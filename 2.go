package main

import "fmt"

func bintang(n int) {
	if n > 0 {
		fmt.Print("*")
		bintang(n - 1)
	}
}

func pola(baris int, jadi int) {
	if jadi <= baris {
		bintang(jadi)
		fmt.Println()
		pola(baris, jadi+1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&n)

	pola(n, 1)
}
