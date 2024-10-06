package main

import (
    "fmt"
    "math"
)

const pi = 3.1415926535

func main() {
    var jejari float64

    fmt.Print("Jejari = ")
    fmt.Scan(&jejari)

    volume := (4.0 / 3.0) * pi * math.Pow(jejari, 3)
    luasPermukaan := 4 * pi * math.Pow(jejari, 2)

    fmt.Printf("Bola dengan jejari %.4f memiliki volume %.4f dan luas kulit %.4f\n", 
               jejari, volume, luasPermukaan)
}