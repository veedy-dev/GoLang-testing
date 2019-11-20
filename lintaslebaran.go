package main

import "fmt"

const MAXSIZE = 99999999

type Trafik [MAXSIZE]int

var (
	arr Trafik
	n   int
)

func catatLaluLintas(arr *Trafik, n int) {
	for i := 1; i <= n; i++ {
		fmt.Scan(&arr[i])
	}
}

func kendaraanTerbanyak(arr *Trafik, n int) int {
	var max, idx int

	for i, a := range arr {
		if a > max {
			max = a
			idx = i
		}
	}
	fmt.Println("\n tipe kendaraan:", idx)
	return max
}
func main() {
	catatLaluLintas(&arr, 5)

	fmt.Println("Kendaraaan terbanyak :", kendaraanTerbanyak(&arr, 5))
}
