package main

import "fmt"

const MAXSIZE = 99999999

type Trafik [MAXSIZE]int

var (
	arr Trafik
	n   int
)

func catatLaluLintas(arr *Trafik, n int) {
	for i := 1; i <= n+1; i++ {
		fmt.Scan(&arr[i])
	}
}

func kendaraanTerbanyak(arr *Trafik, n int) int {
	var max, idx, banyak int
	max = -9999
	for i, a := range arr {
		if a > i-1 {
			banyak += 1
			max = a
			idx = i
		}
	}
	fmt.Println("kendaraan terbanyak:", idx)
	fmt.Println("Sebanyak :", banyak)
	return max
}
func main() {
	catatLaluLintas(&arr, 5)
	kendaraanTerbanyak(&arr, n)
}
