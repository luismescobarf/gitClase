//Crear y cargar una matriz de 4 filas por 4 columnas. Imprimir la diagonal principal.
package main

import "fmt"

func main() {
    var mat [4][4]int
    cont :=0
    for i :=0 ; i < 4; i++ {
        for j := 0; j < 4; j++ {
            fmt.Println("Ingrese el número: ")
            fmt.Scan(&mat[i][j])
        }
    }
    fmt.Println("Matriz completa")
	for i := 0; i < 4; i++{
		
		for j := 0; j < 4; j++ {
            
            if i==j{
                fmt.Print("[",mat[i][i],"]")
                cont++
            }else{
                fmt.Print(mat[i][j]," ")
                cont++
            }
            if cont==4 {
                fmt.Print("\n")
                cont = 0
            }
            
        }
	}

    // fmt.Println(mat)    
    fmt.Println("Elementos de la diagonal principal")
    for i := 0; i < 4; i++ {
        fmt.Print(mat[i][i]," ")
    }    
}