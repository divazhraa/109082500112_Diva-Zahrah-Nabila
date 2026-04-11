package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fact(n, 1)
}
func fact(n, i int) {
	if i > n {
		return
	} else {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
		fact(n, i+1)
	}
}
