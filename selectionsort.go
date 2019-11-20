package main

func main() {
	var (
		n int
		a []int
	)
	i := n - 1
	for i > 0 {
		max := 0
		j := 1
		for j < n {
			if a[j] > a[max] {
				max = j
			}
			j = j + 1
		}
		t := a[max]
		a[max] = a[i]
		a[i] = t
		i = i - 1
	}
}
