package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin *float64, bMax *float64) {
	var i int

	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	i = 1

	for i < n {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}

		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}

		i = i + 1
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var i int
	var total float64

	total = 0

	for i < n {
		total = total + arrBerat[i]
		i = i + 1
	}

	return total / float64(n)
}

func main() {
	var berat arrBalita
	var n, i int
	var min, max, rata float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Print("Masukan berat balita ke-", i+1, ": ")
		fmt.Scan(&berat[i])
	}

	hitungMinMax(berat, n, &min, &max)
	rata = rerata(berat, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}
