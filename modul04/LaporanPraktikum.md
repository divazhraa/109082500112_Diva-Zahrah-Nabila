# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Diva Zahrah Nabila] - [109082500112]</p>

## Unguided 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, isenguntuk mengimplementasikannya ke dalam suatu program.

```go
package main

import "fmt"

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutasi(n, r int, hasil *int) {
	var fn, fr, fnr int
	factorial(n, &fn)
	factorial(r, &fr)
	factorial(n-r, &fnr)
	*hasil = fn / fnr
}

func combinasi(n, r int, hasil *int) {
	var fn, fr, fnr int
	factorial(n, &fn)
	factorial(r, &fr)
	factorial(n-r, &fnr)
	*hasil = fn / (fr * fnr)
}

func main() {
	var a, b, c, d, p, k int
	fmt.Scan(&a, &b, &c, &d)
	permutasi(a, c, &p)
	combinasi(a, c, &k)
	fmt.Println(p, k)

	permutasi(b, d, &p)
	combinasi(b, d, &k)
	fmt.Println(p, k)
}

```
### Output Unguided :

##### Output 
[Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul03/output/output-soal1.png)
[Program ini bertujuan untuk menghitung nilai permutasi dan kombinasi dari dua pasang bilangan dengan menggunakan prosedur. Pengguna memasukkan empat bilangan, lalu program memproses dua pasangan bilangan tersebut dan menampilkan hasil permutasi serta kombinasi masing-masing.Semua hasil dikembalikan melalui parameter pointer sehingga menunjukkan penggunaan prosedur]


### 2. Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur.

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul03/output/output-soal2.png)
[Program ini digunakan untuk menentukan pemenang berdasarkan jumlah soal yang berhasil diselesaikan dan total waktu tercepat. Setiap peserta memasukkan nama dan waktu pengerjaan untuk 8 soal, di mana hanya waktu ≤300 menit yang dihitung. Pemenang dipilih dari peserta dengan soal terbanyak, dan jika sama, maka yang memiliki total waktu paling kecil akan menjadi pemenang.]

### 3. Skiena dan Revilla dalam Programming Challenges mendefinisikan sebuah deret bilangan Deret dimulai dengan sebuah bilangan bulat n. Jika bilangan n saat itu genap, maka suku berikutnya adalah 1⁄2n, tetapi jika ganjil maka suku berikutnya bernilai 3n+1. Rumus yang sama digunakan terus menerus untuk mencari suku berikutnya. Deret berakhir ketika suku terakhir bernilai 1.

```go
package main

import "fmt"

func cetakDeret(n int) {
	for {
		fmt.Print(n, " ")
		if n == 1 {
			break
		}
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	cetakDeret(n)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul03/output/output-soal3.png)
[Program ini berfungsi untuk menampilkan deret angka berdasarkan suatu bilangan input dengan menggunakan aturan tertentu (mirip dengan konsep Collatz). Setelah pengguna memasukkan sebuah bilangan, program akan mencetak nilai tersebut secara berurutan, di mana jika bilangan bernilai genap maka akan dibagi 2, sedangkan jika ganjil akan dikalikan 3 lalu ditambah 1. Proses ini dilakukan berulang hingga mencapai angka 1, dan seluruh deret angka ditampilkan secara berurutan sebagai output.]