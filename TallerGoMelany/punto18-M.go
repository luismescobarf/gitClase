package main

import "fmt"


func menorValor() {
    var v1, v2, v3 int
    fmt.Print("Ingrese primer valor:")
    fmt.Scan(&v1)
    fmt.Print("Ingrese segundo valor:")
    fmt.Scan(&v2)
    fmt.Print("Ingrese tercer valor:")
    fmt.Scan(&v3)
    if v1 < v2 && v1 < v3 {
        fmt.Println("Menor valor:", v1)
    } else {
        if v2 < v3 {
            fmt.Println("Menor valor:", v2)
        } else {
            fmt.Println("Menor valor:", v3)
        }
    }
}


func main() {
    menorValor()
    
}
