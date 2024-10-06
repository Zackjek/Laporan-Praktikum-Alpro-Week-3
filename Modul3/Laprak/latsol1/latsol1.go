package main

import (
	"fmt"
	"math"
)

func calculateX(fx float64) float64 {
	// f(x) = 2/(x+5) + 5
	// fx = 2/(x+5) + 5
	// fx - 5 = 2/(x+5)
	// (fx - 5)(x+5) = 2
	// fx*x + 5*fx - 5*x - 25 = 2
	// fx*x - 5*x + 5*fx - 27 = 0
	// x(fx - 5) = 27 - 5*fx
	// x = (27 - 5*fx) / (fx - 5)
	
	return (27 - 5*fx) / (fx - 5)
}

func main() {
	var fx float64

	fmt.Print("Masukkan nilai f(x): ")
	_, err := fmt.Scanf("%f", &fx)
	if err != nil {
		fmt.Println("Error: Masukan tidak valid")
		return
	}

	if fx == 5 {
		fmt.Println("Error: Nilai f(x) tidak boleh 5 karena akan menyebabkan pembagian dengan nol")
		return
	}

	x := calculateX(fx)
	
	// Pembulatan ke bilangan bulat terdekat
	roundedX := math.Round(x)

	fmt.Printf("Nilai x: %.2f\n", x)
	fmt.Printf("Nilai x (dibulatkan): %.0f\n", roundedX)
}
