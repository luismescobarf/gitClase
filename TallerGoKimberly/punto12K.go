//Crear una matriz de 2 filas y 5 columnas. Realizar la carga de componentes por columna 
//(es decir, primero ingresar toda la primera columna, luego la segunda columna y así sucesivamente) Imprimir luego la matriz
package main

import "fmt"

func main() {
    var matriz [2][5]int
    cont := 0
    for i :=0 ; i < 5; i++ {
        for j := 0; j < 2; j++ {
            fmt.Println("Ingrese el número de la fila ",j," columna ",i)
            fmt.Scan(&matriz[j][i])
        }
    }
    

	 fmt.Println("\nMatriz completa ")
	 for i := 0; i < 2; i++{
	 	for j := 0; j < 5; j++ {
            
            fmt.Print(matriz[i][j]," ")
            cont++
            
            if cont==5 {
                fmt.Print("\n")
                cont = 0
            }
        }
	}
}