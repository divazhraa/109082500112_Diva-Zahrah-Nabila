package main

import "fmt"

const MAX = 1000

type arrInt [MAX]int

func selectionSortAscending(T *arrInt, n int) {
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
func selectionSortDescending(T *arrInt, n int) {
	var t, i, j, idxMax int
	i = 1
	for i <= n-1 {
		idxMax = i - 1
		j = i
		for j < n {
			if T[idxMax] < T[j] {
				idxMax = j
			}
			j++
		}
		t = T[idxMax]
		T[idxMax] = T[i-1]
		T[i-1] = t
		i++
	}
}
func main() {
	var nDaerah, m, i, j, rumah, nGanjil, nGenap int
	var ganjil, genap arrInt
	fmt.Scan(&nDaerah)
	for i = 0; i < nDaerah; i++ {
		nGanjil = 0
		nGenap = 0
		fmt.Scan(&m)
		for j = 0; j < m; j++ {
			fmt.Scan(&rumah)
			if rumah%2 == 1 {
				ganjil[nGanjil] = rumah
				nGanjil++
			} else {
				genap[nGenap] = rumah
				nGenap++
			}
		}
		selectionSortAscending(&ganjil, nGanjil)
		selectionSortDescending(&genap, nGenap)
		for j = 0; j < nGanjil; j++ {
			fmt.Print(ganjil[j])
			if j != nGanjil-1 || nGenap > 0 {
				fmt.Print(" ")
			}
		}
		for j = 0; j < nGenap; j++ {
			fmt.Print(genap[j])

			if j != nGenap-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
