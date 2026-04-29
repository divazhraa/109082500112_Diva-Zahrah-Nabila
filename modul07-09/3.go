package main

import "fmt"

const max = 100

type tabel [max]string

func main() {

	var klubA, klubB string
	var hasil tabel
	var skorA, skorB, i, n int

	fmt.Print("klubA: ")
	fmt.Scan(&klubA)
	fmt.Print("klubB: ")
	fmt.Scan(&klubB)
	i = 0
	for {
		fmt.Print("Pertandingan ", i+1, " : ")
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil[i] = klubA
		} else if skorB > skorA {
			hasil[i] = klubB
		} else {
			hasil[i] = "Draw"
		}
		i++
	}
	n = i
	for i = 0; i < n; i++ {
		fmt.Println("hasil", i+1, " : ", hasil[i])
	}
	fmt.Println("PERTANDINGAN SELESAI")
}
