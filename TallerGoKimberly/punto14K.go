//Definir una matriz de 2 filas y 5 columnas. Realizar su carga e impresión.
// Intercambiar los elementos de la primera fila con la segunda y volver a imprimir la matriz.
package main

import "fmt"

func main() {
    var matriz [2][5]int
    cont := 0
	aux := 0

    for i :=0 ; i < 2; i++ {
        for j := 0; j < 5; j++ {
            fmt.Println("Ingrese el número: ")
            fmt.Scan(&matriz[i][j])
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

   for i := 0; i < 5; i++{
	   	aux = matriz[0][i]
        matriz[0][i] = matriz[1][i]
        matriz[1][i] = aux
    }

   fmt.Println("\nMatriz cambiada ")
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