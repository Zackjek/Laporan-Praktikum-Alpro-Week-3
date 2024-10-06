package main

import "fmt"

func isCumLaude(semesters int, eprtScore int) bool {
    return semesters <= 8 && eprtScore >= 500
}

func main() {
    var semesters, eprtScore int
    fmt.Print("Masukkan jumlah semester dan skor EPrT: ")
    fmt.Scanf("%d %d", &semesters, &eprtScore)
    
    if isCumLaude(semesters, eprtScore) {
        fmt.Println(true)
    } else {
        fmt.Println(false)
    }
}
