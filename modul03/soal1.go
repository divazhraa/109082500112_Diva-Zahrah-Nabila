package main

import "fmt"

func factorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}
func permutasi(n, r int) int {
	return factorial(n) / factorial(n-r)
}
func combinasi(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}
func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	fmt.Println(permutasi(a, c), combinasi(a, c))
	fmt.Println(permutasi(b, d), combinasi(b, d))
}
