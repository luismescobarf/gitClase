package main

import "fmt"

func main() {
    var altura [5]float32
    var suma float32
    for i := 0; i < 5; i++ {
        fmt.Print("Ingrese la altura de la persona :")
        fmt.Scan(&altura[i])
        suma = suma + altura[i]
    }
    promedio := suma / 5
    fmt.Println("Promedio de las alturas:", promedio)
    may := 0
    men := 0
    for i := 0; i < 5; i++ {
        if altura[i] > promedio {
            may++
        } else {
            if altura[i] < promedio {
                men++
            }
        }        
    }
    fmt.Println("Cantidad de personas mayores al promedio:", may)
    fmt.Println("Cantidad de personas menores al promedio:", men)    
}
