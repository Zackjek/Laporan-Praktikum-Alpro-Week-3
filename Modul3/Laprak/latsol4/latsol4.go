package main

import (
	"fmt"
)

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func celsiusToReamur(celsius float64) float64 {
	return celsius * 4 / 5
}

func celsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

func main() {
	var celsius float64

	fmt.Print("Temperatur Celsius: ")
	_, err := fmt.Scanf("%f", &celsius)
	if err != nil {
		fmt.Println("Error membaca input:", err)
		return
	}

	fahrenheit := celsiusToFahrenheit(celsius)
	reamur := celsiusToReamur(celsius)
	kelvin := celsiusToKelvin(celsius)

	fmt.Printf("Derajat Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Derajat Reamur: %.2f\n", reamur)
	fmt.Printf("Derajat Kelvin: %.2f\n", kelvin)
}
