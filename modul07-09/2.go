package main

import (
	"fmt"
	"math"
)

const max int = 100

type tabel [max]int

func isiArray(t *tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i])
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(t[i], " ")
	}
	fmt.Println()
}

func indeksGanjil(t tabel, n int) {
	for i := 1; i < n; i += 2 {
		fmt.Print(t[i], " ")
	}
	fmt.Println()
}

func indeksGenap(t tabel, n int) {
	for i := 0; i < n; i += 2 {
		fmt.Print(t[i], " ")
	}
	fmt.Println()
}
func kelipatan(t tabel, n, x int) {
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(t[i], " ")
		}
	}
	fmt.Println()
}
func hapus(t *tabel, n *int, idx int) {
	for i := idx; i < *n-1; i++ {
		t[i] = t[i+1]
	}
	*n = *n - 1
}

func rataRata(t tabel, n int) float64 {
	var jumlah int
	for i := 0; i < n; i++ {
		jumlah += t[i]
	}
	return float64(jumlah) / float64(n)
}

func standarDeviasi(t tabel, n int) float64 {
	var jumlah float64
	var rata float64

	rata = rataRata(t, n)

	for i := 0; i < n; i++ {
		jumlah += math.Pow(float64(t[i])-rata, 2)
	}
	return math.Sqrt(jumlah / float64(n))
}
func frekuensi(t tabel, n, x int) int {
	var f int
	for i := 0; i < n; i++ {
		if t[i] == x {
			f++
		}
	}
	return f
}

func main() {
	var tab tabel
	var n, x, idx int

	fmt.Scan(&n)
	isiArray(&tab, n)
	cetakArray(tab, n)
	indeksGanjil(tab, n)
	indeksGenap(tab, n)
	fmt.Scan(&x)
	kelipatan(tab, n, x)
	fmt.Scan(&idx)
	hapus(&tab, &n, idx)
	fmt.Println(rataRata(tab, n))
	fmt.Println(standarDeviasi(tab, n))
	fmt.Scan(&x)
	fmt.Println(frekuensi(tab, n, x))
}
