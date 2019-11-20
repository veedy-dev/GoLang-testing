package main

func main() {
	var (
		j, t int
		a    []int
	)
	i := 1
	for i < n {
		t = i - 1
		for j >= 0 && a[j] > t {
			a[j+1] = a[j]
			j = j - 1
		}
		a[j+1] = t
		i = i + 1
	}
}
