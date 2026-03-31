package main

import "fmt"

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutasi(n, r int, hasil *int) {
	var fn, fr, fnr int
	factorial(n, &fn)
	factorial(r, &fr)
	factorial(n-r, &fnr)
	*hasil = fn / fnr
}

func combinasi(n, r int, hasil *int) {
	var fn, fr, fnr int
	factorial(n, &fn)
	factorial(r, &fr)
	factorial(n-r, &fnr)
	*hasil = fn / (fr * fnr)
}

func main() {
	var a, b, c, d, p, k int
	fmt.Scan(&a, &b, &c, &d)
	permutasi(a, c, &p)
	combinasi(a, c, &k)
	fmt.Println(p, k)

	permutasi(b, d, &p)
	combinasi(b, d, &k)
	fmt.Println(p, k)
}
