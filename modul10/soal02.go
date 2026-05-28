package main

import "fmt"

type array [1000]float64

func main() {
	var ikan array
	var x, y int
	var i, j int
	var totalWadah float64
	var rata float64

	fmt.Scan(&x, &y)
	for i = 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	i = 0
	for i < x {
		totalWadah = 0
		j = 0
		for j < y && i < x {
			totalWadah = totalWadah + ikan[i]
			i = i + 1
			j = j + 1
		}
		fmt.Printf("%.2f ", totalWadah)
		rata = rata + totalWadah
	}
	rata = rata / float64((x+y-1)/y)
	fmt.Println()
	fmt.Printf("%.2f\n", rata)
}
