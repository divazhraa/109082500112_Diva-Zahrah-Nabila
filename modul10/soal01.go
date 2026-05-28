package main

import "fmt"

type kelinci [1000]float64

func MinMax(A kelinci, n int, min *float64, max *float64) {
	var j int
	*min = A[0]
	*max = A[0]

	j = 1
	for j < n {
		if A[j] < *min {
			*min = A[j]
		}

		if A[j] > *max {
			*max = A[j]
		}

		j = j + 1
	}
}

func main() {
	var A kelinci
	var n, i int
	var min, max float64
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	MinMax(A, n, &min, &max)
	fmt.Printf("%.1f\n", min)
	fmt.Printf("%.1f\n", max)
}
