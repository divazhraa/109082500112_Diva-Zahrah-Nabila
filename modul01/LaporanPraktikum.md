# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Diva Zahrah Nabila] - [109082500112]</p>

## Unguided 

### 1. [Soal 1]
#### soal01.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul01/output/output-soal01.png)
[Program tersebut digunakan untuk menerima tiga input string dari pengguna lalu menukar urutannya. Program menyimpan input ke dalam variabel satu, dua, tiga, lalu menampilkan output awal sesuai urutan yang dimasukan. Program menggunakan variabel temp sebagai penyimpan sementara untuk menukar posisi nilai sehingga urutannya berubah, yaitu satu menjadi dua, dua menjadi tiga dan tiga menjadi satu.]


### 2. [Soal 2]
#### soal01.go

```go
package main

import "fmt"

func main() {
	var gelas1, gelas2, gelas3, gelas4 string
	berhasil := true

	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&gelas1, &gelas2, &gelas3, &gelas4)

		if !(gelas1 == "merah" && gelas2 == "kuning" && gelas3 == "hijau" && gelas4 == "ungu") {
			berhasil = false
		}
	}
	fmt.Println("BERHASIL:", berhasil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul01/output/output-soal02.png)
[Program tersebut digunakan untuk mengecek urutan warna pada empat gelas. Program melakukan 5 kali percobaan dengan meminta pengguna memasukan empat warna. Jika urutannya merah, kuning, hijau, ungu, maka hasil yang keluar adalah BERHASIL, jika tidak sesuai dengan urutan maka akan FALSE. ]


### 3. [Soal 3]
#### soal01.go

```go
package main

import "fmt"

func main() {

	var berat, kg, sisa, biaya1, biaya2, total int

	fmt.Print("Berat parsel(gram): ")
	fmt.Scan(&berat)

	kg = berat / 1000
	sisa = berat % 1000

	biaya1 = kg * 10000

	if kg > 10 {
		biaya2 = 0
	} else {
		if sisa >= 500 {
			biaya2 = sisa * 5
		} else {
			biaya2 = sisa * 15
		}
	}

	total = biaya1 + biaya2

	fmt.Println("Detail berat:", kg, "kg +", sisa, "gr")
	fmt.Println("Detail biaya: Rp.", biaya1, "+ Rp.", biaya2)
	fmt.Println("Total biaya: Rp.", total)

}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul01/output/output-soal03.png)
[Program tersebut digunakan untuk menghitung biaya pengiriman parsel berdasarkan beratnya. Pengguna memasukan berat dalam gram, lalu program mengubahnya menjadi kg dan sisa gram. Biaya dihitung 10.000 per kg, sedangkan sisa gram dikenakan biaya tambahan sesuai kondisi. Lalu program menampilkan detail berat dan total biaya pengiriman.]