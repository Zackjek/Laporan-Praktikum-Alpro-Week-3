package main

import (
	"fmt"
)

func isDecreasing(number int) bool {
	hundreds := number / 100
	tens := (number / 10) % 10
	ones := number % 10

	return hundreds > tens && tens > ones
}

func main() {
	var number int
	fmt.Print("Masukkan bilangan bulat positif dengan tiga digit: ")
	fmt.Scan(&number)

	if isDecreasing(number) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
