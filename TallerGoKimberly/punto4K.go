//Realizar un programa que pida la carga de dos arreglos numéricos enteros de 4
//elementos. Obtener la suma de los dos arreglos, dicho resultado guardarlo en un
//tercer arreglo del mismo tamaño. Sumar componente a componente.

package main

import "fmt"

func main() {
    var array1 [4]int
    var array2 [4]int
    var arraySuma [4]int

    fmt.Println("Llenar el primer arreglo.")
    for i := 0; i < len(array1); i++ {
        fmt.Println("Ingrese número: ")
        fmt.Scan(&array1[i])
    }
    fmt.Println("Llenar el segundo arreglo.")
    for i := 0; i < len(array2); i++ {
        fmt.Println("Ingrese número:")
        fmt.Scan(&array2[i])
    }
    for i := 0; i < len(arraySuma); i++ {
        arraySuma[i] = array1[i] + array2[i]
    }
    fmt.Println("Arreglo resultante")
    fmt.Println(arraySuma)
}