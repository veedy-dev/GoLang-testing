package main

func main() {
	var n, v int
	var a []int
	kr := 0
	kn := n
	found := false
	for kr < kn && !found {
		med := (kr + kn) / 2
		if a[med] < v {
			kr = med + 1
		} else if a[med] > v {
			kn = med
		} else {
			found = true
		}
	}
}
