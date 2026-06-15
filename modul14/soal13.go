package main

import "fmt"

const MAX = 1000000

type arrInt [MAX]int

func insertionSortAscending(T *arrInt, n int) {
	var temp, i, j int
	i = 1
	for i <= n-1 {
		j = i
		temp = T[j]
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j--
		}
		T[j] = temp
		i++
	}
}
func main() {
	var data arrInt
	var n, bilangan, median int
	n = 0
	for {
		fmt.Scan(&bilangan)
		if bilangan == -5313 {
			break
		}
		if bilangan == 0 {
			insertionSortAscending(&data, n)
			if n%2 == 1 {
				median = data[n/2]
			} else {
				median = (data[(n/2)-1] + data[n/2]) / 2
			}
			fmt.Println(median)
		} else {
			data[n] = bilangan
			n++
		}
	}
}
