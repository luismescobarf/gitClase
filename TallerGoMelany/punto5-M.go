package main

import "fmt"

func main() {
    var cA [5]int
    var cB [5]int
    suma1 := 0
    fmt.Println("Notas del grupo A")
    for i := 0; i < len(cA); i++ {
        fmt.Print("Ingrese nota: ")
        fmt.Scan(&cA[i])
        suma1 = suma1 + cA[i]
    }
    suma2 := 0
    fmt.Println("Notas del grupo B")
    for i := 0; i < len(cB); i++ {
        fmt.Print("Ingrese nota: ")
        fmt.Scan(&cB[i])
        suma2 = suma2 + cB[i]
    }
    promeA := suma1 / 5
    promeB := suma2 / 5
    fmt.Println("Promedio curso A:", promeA)
    fmt.Println("Promedio curso B:", promeB)
    if promeA>promB {
        fmt.Println("El promedio del curso A es mayor ")
    } else {
        fmt.Println("El promedio del curso B es mayor ")
    }    
}