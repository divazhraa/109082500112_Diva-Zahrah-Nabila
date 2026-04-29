# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Diva Zahrah Nabila] - [109082500112]</p>

## Unguided 

### 1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya.
```go
package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}
type lingkaran struct {
	pusat  titik
	radius int
}

func jarak(p, q titik) float64 {
	var d float64
	d = math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
	return d
}
func didalam(c lingkaran, p titik) bool {
	return jarak(c.pusat, p) <= float64(c.radius)
}
func main() {
	var l1, l2 lingkaran
	var p titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.radius)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.radius)
	fmt.Scan(&p.x, &p.y)

	if didalam(l1, p) && didalam(l2, p) {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if didalam(l1, p) {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if didalam(l2, p) {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik diluar lingkaran 1 dan 2")
	}
}


```
### Output Unguided :

##### Output 
[Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul07-09/output/1lingkaran.png)
[Program ini digunakan untuk mengecek posisi sebuah titik terhadap dua lingkaran. Program menyimpan data titik berupa koordinat x dan y, serta data lingkaran berupa titik pusat dan jari-jari. Fungsi 'jarak' dipakai untuk menghitung jarak antara dua titik, sedangkan fungsi 'didalam' digunakan untuk menentukan apakah titik berada di dalam lingkaran. Setelah pengguna memasukkan data dua lingkaran dan satu titik, program akan menampilkan apakah titik berada di dalam kedua lingkaran, salah satu lingkaran, atau di luar keduanya.]


### 2. Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu.

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul07-09/output/2array.png)
[Program ini digunakan untuk mengolah data array bilangan bulat. Program dapat mengisi dan menampilkan isi array, menampilkan elemen pada indeks ganjil dan genap, menampilkan elemen pada indeks kelipatan x, serta menghapus elemen pada indeks tertentu. Selain itu, program juga menghitung nilai rata-rata, standar deviasi, dan frekuensi kemunculan suatu angka dalam array. Semua proses dilakukan berdasarkan input dari pengguna.]

### 3. Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga. Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja. Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.

```go
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


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul07-09/output/3pertandingan.png)
[Program ini berfungsi untuk menampilkan faktor bilangan dari bilangan yang dimasukan. Program menggunakan fungsi rekursif dengan mengecek bilangan mulai dari 1 sampai bilangan yang dimasukkan. Jika ada angka yang habis membagi bilangan inputan maka angka tersebut merupakan bilangan faktor, lalu bilangan itu ditampilkan. Proses akan berhenti saat nilai pemeriksaan melebihi bilangan masukkan.]

### 4. Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom.

```go
package main

import "fmt"

const max int = 127

type tabel [max]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0

	fmt.Print("Teks: ")
	fmt.Scanf("%c", &ch)

	for ch != '.' && *n < max {
		t[*n] = ch
		*n++
		fmt.Scanf("%c", &ch)
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	var temp rune
	for i := 0; i < n/2; i++ {
		temp = t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	var i int
	var sama bool

	sama = true
	i = 0

	for i < n/2 && sama {
		if t[i] != t[n-1-i] {
			sama = false
		}
		i++
	}
	return sama
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	fmt.Print("Palindrom? ", palindrom(tab, m))
	fmt.Println()

	balikanArray(&tab, m)

	fmt.Print("Reverse teks : ")
	cetakArray(tab, m)
}



```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul07-09/output/4reverse.png)
[Program ini digunakan untuk mengolah teks yang dimasukkan hingga tanda titik (.). Teks disimpan dalam array bertipe rune, lalu program mengecek apakah teks tersebut termasuk palindrom, yaitu jika dibaca dari depan dan belakang tetap sama. Setelah itu, program membalik urutan karakter dalam array dan menampilkan hasil teks yang sudah dibalik. ]

