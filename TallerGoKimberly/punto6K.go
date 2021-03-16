//Cargar un arreglo de 10 elementos y verificar posteriormente si el mismo está ordenado de menor a mayor.
package main

import "fmt"

func main() {
    var array [10]int
	orden := 1

    for i := 0; i < len(array); i++ {
        fmt.Println("Ingrese número:")
        fmt.Scan(&array[i])
    }
    for i := 0; i < len(array) - 1; i++ {
        if array[i+1] < array[i] {
            orden=0;
        }
    }
    if orden==1 {
        fmt.Println("El arreglo esta ordenado de menor a mayor")
    } else {
        fmt.Println("El arreglo NO esta ordenado de menor a mayor")
    }    
}	