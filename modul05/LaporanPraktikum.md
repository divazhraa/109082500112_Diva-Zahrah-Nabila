# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Diva Zahrah Nabila] - [109082500112]</p>

## Unguided 

### 1. Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan 𝑆𝑛 = 𝑆𝑛−1 + 𝑆𝑛−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.
```go
package main

import "fmt"

func fibo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return (fibo(n - 1)) + fibo(n-2)
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		fmt.Print(fibo(i), " ")
	}
}


```
### Output Unguided :

##### Output 
[Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul05/output/output1.png)
[Program fibonacci ini bertujuan untuk menghitung penjumlahan dari dua angka sebelumnya, mulai dari 0 dan 1. Program ini menggunakan fungsi rekursif dimana program akan memanggil dirinya sendiri, untuk menghitung nilai fibonacci. Pertama-tama menentukan kondisi berhenti(base-case), yaitu ketika n bernilai 0 atau 1, jika n tidak sama dengan 0 atau 1 maka akan memanggil fungsi yang sama dengan nilai yang lebih kecil(n-1 dan n-2)]


### 2. KBuatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan   menggunakan fungsi rekursif. N adalah masukan dari user.

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	bintang(1, n)
}
func bintang(i, n int) {
	if i > n {
		return
	} else {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
		bintang(i+1, n)
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul05/output/output2.png)
[Program ini berfungsi untuk menampilkan pola bintang berbentuk segitiga sesuai dengan jumlah baris(n) yang dimasukan. Program bekerja menggunakan fungsi rekursif untuk mencetak bintang pada setiap baris hingga batas maksimum yang ditentukan.]

### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fact(n, 1)
}
func fact(n, i int) {
	if i > n {
		return
	} else {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
		fact(n, i+1)
	}
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul05/output/output3.png)
[Program ini berfungsi untuk menampilkan faktor bilangan dari bilangan yang dimasukan. Program menggunakan fungsi rekursif dengan mengecek bilangan mulai dari 1 sampai bilangan yang dimasukkan. Jika ada angka yang habis membagi bilangan inputan maka angka tersebut merupakan bilangan faktor, lalu bilangan itu ditampilkan. Proses akan berhenti saat nilai pemeriksaan melebihi bilangan masukkan.]

### 4. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan  tertentu. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan dari N hingga 1 dan kembali ke N.

```go
package main

import "fmt"

func urut(n int) {
	if n == 1 {
		fmt.Print("1 ")
	} else {
		fmt.Print(n, " ")
		urut(n - 1)
		fmt.Print(n, " ")

	}
}
func main() {
	var n int
	fmt.Scan(&n)
	urut(n)
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul05/output/output4.png)
[Program ini berfungsi untuk menampilkan barisan dari angka yang dimulai dari n menurun hingga 1, lalu naik lagi hingga angka n. Program memanfatkan konsep rekursif dengan dua tahap, yaitu turun hingga mencapai 1 sebagai kondisi berhenti,setelah itu naik sampai angka inputan. Pola yang dihasilkan terlihat simetris. ]

### 5. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil.  Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan ganjil dari 1 hingga N.

```go
package main

import "fmt"

func ganjil(n, i int) {
	if i > n {
		return
	} else {
		fmt.Print(i, " ")
		ganjil(n, i+2)
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	ganjil(n, 1)
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul05/output/output5.png)
[Program ini menampilkan deretan bilangan ganjil dari 1 hingga angka yang dimasukkan. Program menggunakan fungsi rekursif dengan dimulai dari angka 1, lalu menambahkan 2 pada setiap pemanggilan untuk mendapatkan bilangan ganjil. Proses akan berhenti ketika angka yang dihasilkan melebihi angka yang dimasukkan pengguna.  ]

### 6. Buatlah program yang mengimplementasikan rekursif untuk mencari hasil pangkat dari dua buah bilangan. Masukan terdiri dari bilangan bulat x dan y. Keluaran terdiri dari hasil x dipangkatkan y. 

```go
package main

import "fmt"

func pangkat(x, y int) int {
	if y == 0 {
		return 1
	} else {
		return x * (pangkat(x, y-1))
	}
}
func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Println(pangkat(x, y))
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul05/output/output6.png)
[Program ini berfungsi untuk menghitung hasil perpangkatan dari dua bilangan yang dimasukkan pengguna, yaitu bilangan(x) dipangkatkan dengan bilangan (y). Program menggunakan fungsi rekursif untuk mengkalikan bilangan dasar berulang sebanyak nilai pangkat. Program akan berhenti jika pangkat bernilai 0, dan hasilnya adalah 1. ]