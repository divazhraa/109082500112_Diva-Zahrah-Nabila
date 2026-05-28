# <h1 align="center">Laporan Praktikum Modul 10 - ... </h1>
<p align="center">[Diva Zahrah Nabila] - [109082500112]</p>

## Unguided 

### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual. Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual. Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan
terbesar.
```go
package main

import "fmt"

type kelinci [1000]float64

func MinMax(A kelinci, n int, min *float64, max *float64) {
	var j int
	*min = A[0]
	*max = A[0]

	j = 1
	for j < n {
		if A[j] < *min {
			*min = A[j]
		}

		if A[j] > *max {
			*max = A[j]
		}

		j = j + 1
	}
}

func main() {
	var A kelinci
	var n, i int
	var min, max float64
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	MinMax(A, n, &min, &max)
	fmt.Printf("%.1f\n", min)
	fmt.Printf("%.1f\n", max)
}



```
### Output Unguided :

##### Output 
[Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul10/output/SOAL1.png)
[Program ini dipakai untuk mencari berat kelinci yang paling kecil dan paling besar. Pengguna memasukkan jumlah kelinci lalu memasukkan berat masing-masing kelinci. Setelah itu program mengecek semua data dan menampilkan berat terkecil serta terbesar]


### 2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual. Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual. Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yangdimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul10/output/SOAL2.png)
[ Program ini dibuat untuk menghitung berat ikan yang dimasukkan ke dalam beberapa wadah. Pengguna terlebih dahulu memasukkan jumlah ikan dan berapa banyak ikan yang bisa dimuat dalam satu wadah. Setelah itu, pengguna memasukkan berat setiap ikan. Program kemudian mengelompokkan ikan ke dalam wadah sesuai urutan input. Berat ikan di setiap wadah dijumlahkan dan hasilnya ditampilkan. Setelah semua wadah selesai dihitung, program juga mencari rata-rata berat dari seluruh wadah dan menampilkannya di akhir.]

### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul10/output/SOAL3.png)
[Program ini digunakan untuk mencatat berat balita di posyandu. Pengguna memasukkan jumlah data balita lalu memasukkan berat masing-masing balita. Setelah data dimasukkan, program akan mencari berat balita paling kecil, paling besar, dan menghitung rata-rata berat balita. Hasil akhirnya kemudian ditampilkan.]

