package main

import "fmt"

type arrInt [1000]int

func insertionSort(A *arrInt, n int) {
	var i, j, temp int
	for i = 1; i < n; i++ {
		temp = A[i]
		j = i
		for j > 0 && temp < A[j-1] {
			A[j] = A[j-1]
			j--
		}
		A[j] = temp
	}
}
func main() {
	var A arrInt
	var n, x int
	n = 0
	for {
		fmt.Scan(&x)
		if x < 0 {
			break
		}
		A[n] = x
		n++
	}
	insertionSort(&A, n)
	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()
	if n <= 1 {
		fmt.Println("Data berjarak 0")
		return
	}
	jarak := A[1] - A[0]
	tetap := true
	for i := 2; i < n; i++ {
		if A[i]-A[i-1] != jarak {
			tetap = false
			break
		}
	}
	if tetap {
		fmt.Println("Data berjarak", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
