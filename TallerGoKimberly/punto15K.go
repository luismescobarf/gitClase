//Definir un slice vacío. En una estructura repetitiva ingresar enteros y añadirlos al slice. 
//Cuando se ingrese el número -1 finalizar la carga y mostrar cuantos números ingresados son mayor al promedio de los valores ingresados
package main

import "fmt"

func main() {
    var slice []int
    var valor int
    suma := 0
    numeroMayor := 0
	
	for {
        fmt.Println("Ingrese un número entero (digite -1 para finalizar): ")
        fmt.Scan(&valor)
        if valor==-1 {
            break
        }
        slice = append(slice, valor)
        suma = suma + valor
    }
    promedio := suma / len(slice)
    
    for i := 0; i < len (slice); i++ {
        if slice[i] > promedio {
            numeroMayor++
        }
    }

    fmt.Println("El slice completo es: ")
    fmt.Println(slice)
    fmt.Println("\nPromedio:", promedio)
    fmt.Println("La cantidad de números mayores al promedio son:", numeroMayor)
}