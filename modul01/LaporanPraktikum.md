# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Diva Zahrah Nabila] - [109082500112]</p>

## Unguided 

### 1. Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut? 
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
![Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul01/output/output-soal01.png)
[Program tersebut digunakan untuk menerima tiga input string dari pengguna lalu menukar urutannya. Program menyimpan input ke dalam variabel satu, dua, tiga, lalu menampilkan output awal sesuai urutan yang dimasukan. Program menggunakan variabel temp sebagai penyimpan sementara untuk menukar posisi nilai sehingga urutannya berubah, yaitu satu menjadi dua, dua menjadi tiga dan tiga menjadi satu.]


### 2. Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’, ‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang. Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya.
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
![Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul01/output/output-soal01.png)
[Program tersebut digunakan untuk mengecek urutan warna pada empat gelas. Program melakukan 5 kali percobaan dengan meminta pengguna memasukan empat warna. Jika urutannya merah, kuning, hijau, ungu, maka hasil yang keluar adalah BERHASIL, jika tidak sesuai dengan urutan maka akan FALSE. ]


### 3. PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan sebagai berikut! Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.
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
![Screenshot Output Unguided 1_!] (https://github.com/divazhraa/109082500112_Diva-Zahrah-Nabila/blob/main/modul01/output/output-soal01.png)
[Program tersebut digunakan untuk menghitung biaya pengiriman parsel berdasarkan beratnya. Pengguna memasukan berat dalam gram, lalu program mengubahnya menjadi kg dan sisa gram. Biaya dihitung 10.000 per kg, sedangkan sisa gram dikenakan biaya tambahan sesuai kondisi. Lalu program menampilkan detail berat dan total biaya pengiriman.]