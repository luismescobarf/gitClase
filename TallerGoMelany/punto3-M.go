package main

import "fmt"

func main() {
    var array [8]int
    sumaTotal := 0
    sumaMayor := 0
    cant := 0
    for i := 0; i < len(array); i++ {
        fmt.Print("Ingrese elemento:")
        fmt.Scan(&array[i])
        sumaTotal = sumaTotal + array[i]
        if array[i] > 50 {
            cant++
        }
        if array[i] > 36 {
            sumaMayor = sumaMayor + array[i]
        }        
    }
    fmt.Println("La suma de los 8 elementos es: ", sumaTotal)
    fmt.Println("La suma de los elementos mayores a 36: ", sumaMayor)    
    fmt.Println("La cantidad de valores mayores a 50 es: ", cant)
}