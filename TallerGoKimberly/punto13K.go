//Crear una matriz de 3x4. Realizar la carga y luego imprimir el elemento mayor.
package main

import "fmt"

func main() {
    var matriz [3][4]int
    aux := 0

    for i :=0 ; i < 3; i++ {
        for j := 0; j < 4; j++ {
            fmt.Println("Ingrese el número: ")
            fmt.Scan(&matriz[i][j])
        }
    }
    

	 fmt.Println("\nEl número mayor es: ")
	 for i := 0; i < 3; i++{
	 	for j := 0; j < 4; j++ {
            
            if matriz[i][j]>aux {
				aux=matriz[i][j]
            }
        }
	}
	fmt.Println(aux)
}