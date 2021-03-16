//Desarrollar un programa que permita ingresar un arreglo de 8 elementos, e informe:
// • El valor acumulado de todos los elementos del arreglo.
// • El valor acumulado de los elementos del arreglo que sean mayores a 36.
// • Cantidad de valores mayores a 50.
package main

import "fmt"

func main() {
    var array [8]int
    sumaTotal := 0
    sumaMayor36 := 0
    cantidad := 0
    for i := 0; i < len(array); i++ {
        fmt.Print("Ingrese un número: ")
        fmt.Scan(&array[i])
        sumaTotal = sumaTotal + array[i]
        if array[i] > 50 {
            cantidad++
        }
        if array[i] > 36 {
            sumaMayor36 = sumaMayor36 + array[i]
        }        
    }
    fmt.Println("\nEl valor acumulado de todos los elementos del arreglo es:", sumaTotal)
    fmt.Println("\nEl valor acumulado de los elementos del arreglo que sean mayores a 36 es:", sumaMayor36)    
    fmt.Println("\nCantidad de valores mayores a 50 es:", cantidad)
}