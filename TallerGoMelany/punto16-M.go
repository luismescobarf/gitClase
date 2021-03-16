package main

import "fmt"

func main() {
    var can int
    fmt.Print("Ingrese la cantidad de productos: ")
    fmt.Scan(&can)
    productos := make([]string, can)
    precio := make([]float64, can)
    for i := 0; i < len(productos); i++ {
        fmt.Print("Ingrese nombre del producto: ")
        fmt.Scan(&productos[i])
        fmt.Print("Ingrese el precio del producto:")
        fmt.Scan(&precio[i])
    }
    fmt.Println("Lista de todos los datos ingresados")
    for i := 0; i < len(productos); i++ {
        fmt.Println(productos[i], "-", precio[i])
    }    
}