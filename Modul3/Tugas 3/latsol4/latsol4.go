package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Print("Masukkan bilangan tiga digit: ")
	fmt.Scan(&number)
	// Mendapatkan digit pertama, kedua, dan ketiga
	digit1 := number / 100         // Digit pertama
	digit2 := (number / 10) % 10   // Digit kedua
	digit3 := number % 10          // Digit ketiga
	// Mengecek apakah digit-digit terurut (menaik atau menurun)
	if (digit1 < digit2 && digit2 < digit3) || (digit1 > digit2 && digit2 > digit3) {
		fmt.Println("true")  // Jika terurut
	} else {
		fmt.Println("false") // Jika tidak terurut
	}
}