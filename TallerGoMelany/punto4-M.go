package main

import "fmt"

func main() {
    var vec1 [4]int
    var vec2 [4]int
    var vecSuma[4]int
    fmt.Println("Carga del primer vector")
    for i := 0; i < len(vec1); i++ {
        fmt.Print("Ingrese elemento: ")
        fmt.Scan(&vec1[i])
    }
    fmt.Println("Carga del segundo vector")
    for i := 0; i < len(vec2); i++ {
        fmt.Print("Ingrese elemento: ")
        fmt.Scan(&vec2[i])
    }
    for i := 0; i < len(vecSuma); i++ {
        vecSuma[i] = vec1[i] + vec2[i]
    }
    fmt.Println("Resultado: ")
    fmt.Println(vecSuma)
}