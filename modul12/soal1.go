package main

import "fmt"

type arrSuara [21]int

func bacaSuara(T *arrSuara, masuk, sah *int) {
	var x int

	fmt.Scan(&x)

	for x != 0 {
		*masuk = *masuk + 1

		if x >= 1 && x <= 20 {
			T[x] = T[x] + 1
			*sah = *sah + 1
		}

		fmt.Scan(&x)
	}
}

func cetakSuara(T arrSuara) {
	var i int

	for i = 1; i <= 20; i++ {
		if T[i] > 0 {
			fmt.Println(i, ":", T[i])
		}
	}
}

func main() {
	var T arrSuara
	var masuk, sah int

	bacaSuara(&T, &masuk, &sah)

	fmt.Println("Suara masuk:", masuk)
	fmt.Println("Suara sah:", sah)

	cetakSuara(T)
}
