package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	bintang(1, n)
}
func bintang(i, n int) {
	if i > n {
		return
	} else {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
		bintang(i+1, n)
	}
}
