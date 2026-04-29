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
