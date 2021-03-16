//Ingresar valores enteros por teclado y sumarlos. Cada vez que se carga un valor pedir
//si quiere continuar con la carga de otro dato ingresando el string "si" o "no".
package main

import "fmt"

func main() {
    var valor int
    var opcion string
    suma := 0
    for {
        fmt.Println("Ingrese un número entero: ")
        fmt.Scan(&valor)
        suma = suma + valor
        fmt.Println("Desea cargar otro valor(escriba si o no)?")
        fmt.Scan(&opcion)
        if opcion=="no" {
            break
        }
    }
    fmt.Println("La suma total de los valores ingresados es: ", suma)
}