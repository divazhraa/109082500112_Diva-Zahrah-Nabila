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
