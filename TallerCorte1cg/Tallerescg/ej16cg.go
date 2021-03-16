/*Confeccionar un programa que permita almacenar en dos slices los nombres de
productos y sus precios (hacer que en las componentes de los mismos subíndices se
guarden datos relacionados). Cuando inicia el programa solicitar la cantidad de
productos a cargar. Mostrar los datos de los dos slices.*/
package main

import "fmt"

func main() {
    var cant int
    fmt.Print("Cuantos productos ingresará:")
    fmt.Scan(&cant)
    productos := make([]string, cant)
    precios := make([]float32, cant)
    for f := 0; f < len(productos); f++ {
        fmt.Print("Ingrese producto:")
        fmt.Scan(&productos[f])
        fmt.Print("Ingrese precio:")
        fmt.Scan(&precios[f])
    }
    fmt.Println("Slices")
    for f := 0; f < len(productos); f++ {
        fmt.Println(productos[f], "-", precios[f])
    }    
}