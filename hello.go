package main

import "fmt"

type icecream int

var x icecream
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 66
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
