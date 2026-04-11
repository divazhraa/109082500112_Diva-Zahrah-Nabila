package main

import "fmt"

func ganjil(n, i int) {
	if i > n {
		return
	} else {
		fmt.Print(i, " ")
		ganjil(n, i+2)
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	ganjil(n, 1)
}

// tidak menggunakan modulo karena jika menggunakan modulo angka akan diperiksa satu satu
