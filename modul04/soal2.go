package main

import "fmt"

func hitungScore(soal *int, skor *int) {
	var waktu int
	*skor = 0
	*soal = 0
	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu <= 300 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama, pemenang string
	var soal, skor, minSkor int
	var maxSoal = -1

	for {
		fmt.Scan(&nama)
		if nama == "selesai" {
			break
		}
		hitungScore(&soal, &skor)

		if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			maxSoal = soal
			minSkor = skor
			pemenang = nama
		}
	}
	fmt.Println(pemenang, maxSoal, minSkor)
}
