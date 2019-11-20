package main

import (
	"fmt"
	"regexp"
)

func main() {
	var berat int
	fmt.Scanln(&berat)
	re := regexp.MustCompile("[0-9]+")
	fmt.Println(re.FindAllString("a", -1))
}
