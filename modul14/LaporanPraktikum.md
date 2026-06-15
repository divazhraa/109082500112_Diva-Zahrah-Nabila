# <h1 align="center">Laporan Praktikum Modul 14 - ... </h1>
<p align="center">[Diva Zahrah Nabila] - [109082500112]</p>

## Unguided 

### 1. Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu. Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort.
```go
package main

import "fmt"

const MAX = 1000000

type arrInt [MAX]int

func selectionSort(T *arrInt, n int) {
	var t, i, j, idxMin int
	i = 1
	for i <= n-1 {
		idxMin = i - 1
		j = i
		for j < n {
			if T[idxMin] > T[j] {
				idxMin = j
			}
			j++
		}
		t = T[idxMin]
		T[idxMin] = T[i-1]
		T[i-1] = t
		i++
	}
}

func main() {
	var nDaerah, m, i, j int
	var rumah arrInt
	fmt.Scan(&nDaerah)
	for i = 0; i < nDaerah; i++ {
		fmt.Scan(&m)
		for j = 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}
		selectionSort(&rumah, m)
		for j = 0; j < m; j++ {
			fmt.Print(rumah[j])
			if j != m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

```
### Output Unguided :

##### Output 
[Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul14/output/soal11.png)
[Program Rumah Kerabat digunakan untuk membantu Hercules mengurutkan nomor rumah seluruh kerabatnya yang berada di beberapa daerah. Program menerima inputaan  berupa jumlah daerah, kemudian untuk setiap daerah dimasukkan banyaknya rumah beserta nomor rumah para kerabat. Selanjutnya, nomor rumah pada setiap daerah diurutkan secara membesar (ascending) menggunakan algoritma Selection Sort, yaitu dengan mencari nilai terkecil dari bagian data yang belum terurut dan menukarnya ke posisi yang seharusnya. Setelah proses pengurutan selesai, program menampilkan daftar nomor rumah yang telah tersusun secara terurut pada masing-masing daerah.]


### 2. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil.
```go
package main

import "fmt"

const MAX = 1000

type arrInt [MAX]int

func selectionSortAscending(T *arrInt, n int) {
	var t, i, j, idxMin int
	i = 1
	for i <= n-1 {
		idxMin = i - 1
		j = i
		for j < n {
			if T[idxMin] > T[j] {
				idxMin = j
			}
			j++
		}
		t = T[idxMin]
		T[idxMin] = T[i-1]
		T[i-1] = t
		i++
	}
}
func selectionSortDescending(T *arrInt, n int) {
	var t, i, j, idxMax int
	i = 1
	for i <= n-1 {
		idxMax = i - 1
		j = i
		for j < n {
			if T[idxMax] < T[j] {
				idxMax = j
			}
			j++
		}
		t = T[idxMax]
		T[idxMax] = T[i-1]
		T[i-1] = t
		i++
	}
}
func main() {
	var nDaerah, m, i, j, rumah, nGanjil, nGenap int
	var ganjil, genap arrInt
	fmt.Scan(&nDaerah)
	for i = 0; i < nDaerah; i++ {
		nGanjil = 0
		nGenap = 0
		fmt.Scan(&m)
		for j = 0; j < m; j++ {
			fmt.Scan(&rumah)
			if rumah%2 == 1 {
				ganjil[nGanjil] = rumah
				nGanjil++
			} else {
				genap[nGenap] = rumah
				nGenap++
			}
		}
		selectionSortAscending(&ganjil, nGanjil)
		selectionSortDescending(&genap, nGenap)
		for j = 0; j < nGanjil; j++ {
			fmt.Print(ganjil[j])
			if j != nGanjil-1 || nGenap > 0 {
				fmt.Print(" ")
			}
		}
		for j = 0; j < nGenap; j++ {
			fmt.Print(genap[j])

			if j != nGenap-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul14/output/soal12.png)
[Program menerima masukan berupa beberapa nomor rumah kerabat pada setiap daerah. Nomor rumah ganjil dan genap dipisahkan ke dalam dua kelompok yang berbeda. Nomor rumah ganjil kemudian diurutkan secara membesar (ascending), sedangkan nomor rumah genap diurutkan secara mengecil (descending) menggunakan algoritma Selection Sort. Setelah proses pengurutan selesai, program menampilkan seluruh nomor rumah ganjil terlebih dahulu, kemudian diikuti oleh nomor rumah genap.]

### 3. Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruantinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya?"Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semuadata merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke bawah." Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0.

```go
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
	var n,bilangan,median int
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul14/output/soal13.png)
[Program Median digunakan untuk mencari nilai median dari sekumpulan bilangan bulat yang dimasukkan oleh pengguna. Setiap bilangan positif akan disimpan ke dalam array. Ketika pengguna memasukkan angka 0, program akan mengurutkan data menggunakan Insertion Sort, kemudian menghitung dan menampilkan nilai median dari data tersebut. Jika jumlah data ganjil, median diambil dari nilai tengah. Jika jumlah data genap, median diperoleh dari rata-rata dua nilai tengah. Program akan berhenti ketika menerima input -5313.]

### 2.1 Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya. Masukan terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non negatif saja yang disimpan ke dalam array. Keluaran terdiri dari dua baris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam array. "Data berjarak x" atau "data berjarak tidak tetap".

```go
package main

import "fmt"

type arrInt [1000]int

func insertionSort(A *arrInt, n int) {
	var i, j, temp int
	for i = 1; i < n; i++ {
		temp = A[i]
		j = i
		for j > 0 && temp < A[j-1] {
			A[j] = A[j-1]
			j--
		}
		A[j] = temp
	}
}
func main() {
	var A arrInt
	var n, x int
	n = 0
	for {
		fmt.Scan(&x)
		if x < 0 {
			break
		}
		A[n] = x
		n++
	}
	insertionSort(&A, n)
	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()
	if n <= 1 {
		fmt.Println("Data berjarak 0")
		return
	}
	jarak := A[1] - A[0]
	tetap := true
	for i := 2; i < n; i++ {
		if A[i]-A[i-1] != jarak {
			tetap = false
			break
		}
	}
	if tetap {
		fmt.Println("Data berjarak", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul14/output/soal21.png)
[Program ini membaca beberapa bilangan bulat hingga pengguna memasukkan bilangan negatif. Data kemudian diurutkan menggunakan Insertion Sort dan diperiksa apakah selisih antar angkanya sama. Jika sama, program menampilkan "Data berjarak x", jika tidak maka "Data berjarak tidak tetap". ]

### 2.2 Masukan terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya, masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari.Keluaran terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku yang dicari sesuai rating yang diberikan pada masukan baris terakhir.

```go
package main

import "fmt"

const NMAX = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating int
}

type DaftarBuku [NMAX]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(
			&pustaka[i].id,
			&pustaka[i].judul,
			&pustaka[i].penulis,
			&pustaka[i].penerbit,
			&pustaka[i].eksemplar,
			&pustaka[i].tahun,
			&pustaka[i].rating,
		)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	var idx, i int
	idx = 0
	for i = 1; i < n; i++ {
		if pustaka[i].rating > pustaka[idx].rating {
			idx = i
		}
	}

	fmt.Println("Buku Terfavorit:")
	fmt.Println("Judul    :", pustaka[idx].judul)
	fmt.Println("Penulis  :", pustaka[idx].penulis)
	fmt.Println("Penerbit :", pustaka[idx].penerbit)
	fmt.Println("Tahun    :", pustaka[idx].tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var i, j int
	var temp Buku
	for i = 1; i < n; i++ {
		temp = pustaka[i]
		j = i
		for j > 0 && temp.rating > pustaka[j-1].rating {
			pustaka[j] = pustaka[j-1]
			j--
		}
		pustaka[j] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	var i, batas int
	fmt.Println("5 Buku Dengan Rating Tertinggi:")
	if n < 5 {
		batas = n
	} else {
		batas = 5
	}
	for i = 0; i < batas; i++ {
		fmt.Println(i+1, ".", pustaka[i].judul, "(Rating:", pustaka[i].rating, ")")
	}
}

func CariBuku(pustaka DaftarBuku, n, r int) {
	var kiri, kanan, tengah int
	var ketemu bool
	kiri = 0
	kanan = n - 1
	ketemu = false
	for kiri <= kanan && !ketemu {
		tengah = (kiri + kanan) / 2
		if pustaka[tengah].rating == r {
			fmt.Println("Buku Ditemukan")
			fmt.Println("ID        :", pustaka[tengah].id)
			fmt.Println("Judul     :", pustaka[tengah].judul)
			fmt.Println("Penulis   :", pustaka[tengah].penulis)
			fmt.Println("Penerbit  :", pustaka[tengah].penerbit)
			fmt.Println("Tahun     :", pustaka[tengah].tahun)
			fmt.Println("Eksemplar :", pustaka[tengah].eksemplar)
			fmt.Println("Rating    :", pustaka[tengah].rating)
			ketemu = true
		} else if r > pustaka[tengah].rating {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	if !ketemu {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}
func main() {
	var pustaka DaftarBuku
	var n, ratingCari int
	fmt.Scan(&n)
	DaftarkanBuku(&pustaka, n)
	fmt.Scan(&ratingCari)
	CetakTerfavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)
	CariBuku(pustaka, n, ratingCari)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul14/output/soal22.png)
[ Program ini mengelola data buku yang berisi ID, judul, penulis, penerbit, jumlah eksemplar, tahun terbit, dan rating. Data diurutkan berdasarkan rating menggunakan Insertion Sort, lalu ditampilkan buku dengan rating tertinggi dan 5 buku terbaik. Program juga dapat mencari buku berdasarkan rating menggunakan Binary Search.]
