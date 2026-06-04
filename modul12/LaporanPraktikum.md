# <h1 align="center">Laporan Praktikum Modul 12 - ... </h1>
<p align="center">[Diva Zahrah Nabila] - [109082500112]</p>

## Unguided 

### 1. Pada pemilihan ketua RT yang baru saja berlangsung, terdapat 20 calon ketua yang bertanding memperebutkan suara warga. Perhitungan suara dapat segera dilakukan karena warga cukup mengisi formulir dengan nomor dari calon ketua RT yang dipilihnya. Seperti biasa, selalu ada pengisian yang tidak tepat atau dengan nomor pilihan di luar yang tersedia, sehingga data juga harus divalidasi. Tugas Anda untuk membuat program mencari siapa yang memenangkan pemilihan ketua RT.
```go
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

```
### Output Unguided :

##### Output 
[Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul12/output/SOAL1.png)
[Program ini digunakan untuk membaca data suara pemilihan ketua RT, memvalidasi setiap suara yang masuk, kemudian menghitung jumlah suara masuk, jumlah suara sah, dan banyaknya suara yang diperoleh masing-masing calon. Data dianggap sah apabila bernilai antara 1 sampai 20, dan proses input berakhir saat pengguna memasukkan angka 0]


### 2. Berdasarkan program sebelumnya, buat program pilkart yang mencari siapa pemenang pemilihan ketua RT. Sekaligus juga ditentukan bahwa wakil ketua RT adalah calon yang mendapatkan suara terbanyak kedua. Jika beberapa calon mendapatkan suara terbanyak yang sama, ketua terpilih adalah dengan nomor peserta yang paling kecil dan wakilnya dengan nomor peserta terkecil berikutnya.
```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul12/output/SOAL2.png)
[ Program ini digunakan untuk menghitung hasil pemilihan ketua RT berdasarkan suara yang masuk. Setelah seluruh suara valid dihitung, program akan menentukan calon dengan suara terbanyak sebagai ketua RT dan calon dengan suara terbanyak kedua sebagai wakil ketua RT. Jika terdapat jumlah suara yang sama, dipilih calon dengan nomor yang lebih kecil]

### 3. Diberikan n data integer positif dalam keadaan terurut membesar dan sebuah integer lain k, apakah bilangan k tersebut ada dalam daftar bilangan yang diberikan? Jika ya, berikan indeksnya, jika tidak sebutkan "TIDAK ADA". Masukan terdiri dari dua baris. Baris pertama berisi dua buah integer positif, yaitu n dan k. menyatakan banyaknya data, dimana 1 < n <= 1000000. k adalah bilangan yang ingin dicari. Baris kedua berisi n buah data integer positif yang sudah terurut membesar. Keluaran terdiri dari satu baris saja, yaitu sebuah bilangan yang menyatakan posisi data yang dicari (k) dalam kumpulan data yang diberikan. Posisi data dihitung dimulai dari angka 0. Atau memberikan keluaran "TIDAK ADA" jika data k tersebut tidak ditemukan dalam kumpulan.

```go
package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func isiArray(n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	var found int = -1
	var kr int = 0
	var kn int = n - 1
	var med int

	for kr <= kn && found == -1 {
		med = (kr + kn) / 2

		if k < data[med] {
			kn = med - 1
		} else if k > data[med] {
			kr = med + 1
		} else {
			found = med
		}
	}

	return found
}

func main() {
	var n, k int
	var hasil int

	fmt.Scan(&n, &k)

	isiArray(n)

	hasil = posisi(n, k)

	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul12/output/SOAL3.png)
[Program ini digunakan untuk mencari posisi suatu bilangan pada kumpulan data integer yang sudah terurut membesar. Pencarian dilakukan menggunakan algoritma Binary Search sehingga proses pencarian menjadi lebih cepat. Program akan menampilkan indeks data jika ditemukan, atau menampilkan tulisan "TIDAK ADA" apabila data yang dicari tidak terdapat dalam array]

