package main

import "fmt"

func main() {
    var cursoa [5]int
    var cursob [5]int
    suma1 := 0
    fmt.Println("Carga de notas del curso A")
    for f := 0; f < len(cursoa); f++ {
        fmt.Print("Ingrese nota:")
        fmt.Scan(&cursoa[f])
        suma1 = suma1 + cursoa[f]
    }
    suma2 := 0
    fmt.Println("Carga de notas del curso B")
    for f := 0; f < len(cursob); f++ {
        fmt.Print("Ingrese nota:")
        fmt.Scan(&cursob[f])
        suma2 = suma2 + cursob[f]
    }
    promedioa := suma1 / 5
    promediob := suma2 / 5
    fmt.Println("Promedio curso A:", promedioa)
    fmt.Println("Promedio curso B:", promediob)
    if promedioa>promediob {
        fmt.Println("El curso A tiene un promedio mayor.")
    } else {
        fmt.Println("El curso B tiene un promedio mayor.")
    }    
}
