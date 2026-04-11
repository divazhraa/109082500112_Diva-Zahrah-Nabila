package main

import "fmt"

func fibo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return (fibo(n - 1)) + fibo(n-2)
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		fmt.Print(fibo(i), " ")
	}
}
