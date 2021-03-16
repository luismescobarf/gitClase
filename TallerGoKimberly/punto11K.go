//Crear y cargar una matriz de 3 filas por 4 columnas. Imprimir la primera fila. Imprimir
// la última fila e imprimir la primera columna.
package main

import "fmt"

func main() {
    var matriz [3][4]int
    cont := 0

    for i :=0 ; i < 3; i++ {
        for j := 0; j < 4; j++ {
            fmt.Println("Ingrese el numero: ")
            fmt.Scan(&matriz[i][j])
        }
    }
    
	fmt.Println("\nMatriz completa ")
	for i := 0; i < 3; i++{
		for j := 0; j < 4; j++ {
            
            fmt.Print(matriz[i][j]," ")
            cont++
            
            if cont==4 {
                fmt.Print("\n")
                cont = 0
            }
        }
	}

	fmt.Println("\nPrimera y última fila")
	for i := 0; i < 3; i++{
		for j := 0; j < 4; j++ {
            
            if i==0||i==2{
                fmt.Print(matriz[i][j]," ")
                cont++
            }
            if cont==4 {
                fmt.Print("\n")
                cont = 0
            }
        }
	}

    fmt.Println("\nPrimera columna")

	for i := 0; i < 3; i++{
		for j := 0; j < 4; j++ {
            
            if j==0{
                fmt.Println(matriz[i][j])
                
            }
        }
	}
}