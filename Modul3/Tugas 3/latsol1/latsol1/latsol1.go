package main

import "fmt"

func main() {
    var number int
    fmt.Print("Masukkan bilangan bulat positif: ")
    fmt.Scan(&number)

    if number%2 != 0 {
        fmt.Println(true) // Bilangan ganjil
    } else {
        fmt.Println(false) // Bilangan genap
    }
}
