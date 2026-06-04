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

func ketuaRT(T arrSuara) int {
	var i, ketua int

	ketua = 1

	for i = 2; i <= 20; i++ {
		if T[i] > T[ketua] {
			ketua = i
		}
	}

	return ketua
}

func wakilRT(T arrSuara, ketua int) int {
	var i, wakil int

	if ketua == 1 {
		wakil = 2
	} else {
		wakil = 1
	}

	for i = 1; i <= 20; i++ {
		if i != ketua {
			if T[i] > T[wakil] {
				wakil = i
			}
		}
	}

	return wakil
}

func main() {
	var T arrSuara
	var masuk, sah int
	var ketua, wakil int

	bacaSuara(&T, &masuk, &sah)

	ketua = ketuaRT(T)
	wakil = wakilRT(T, ketua)

	fmt.Println("Suara masuk:", masuk)
	fmt.Println("Suara sah:", sah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}
