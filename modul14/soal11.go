package main

import "fmt"

const MAX = 1000000

type arrInt [MAX]int

func selectionSort(T *arrInt, n int) {
	var t, i, j, idxMin int
	i = 1
	for i <= n-1 {
		idxMin = i - 1
		j = i
		for j < n {
			if T[idxMin] > T[j] {
				idxMin = j
			}
			j++
		}
		t = T[idxMin]
		T[idxMin] = T[i-1]
		T[i-1] = t
		i++
	}
}

func main() {
	var nDaerah, m, i, j int
	var rumah arrInt
	fmt.Scan(&nDaerah)
	for i = 0; i < nDaerah; i++ {
		fmt.Scan(&m)
		for j = 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}
		selectionSort(&rumah, m)
		for j = 0; j < m; j++ {
			fmt.Print(rumah[j])
			if j != m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
