package main

import "fmt"

func main() {
    var valor int
    var opcion string
    suma := 0
    for {
        fmt.Print("Ingrese un numero entero: ")
        fmt.Scan(&valor)
        suma = suma + valor
        fmt.Print("¿Desea sumar otro numero (si / no)?")
        fmt.Scan(&opcion)
        if opcion=="no" {
            break
        }
    }
    fmt.Print("La suma de los numeros ingresados es ", suma)
}

