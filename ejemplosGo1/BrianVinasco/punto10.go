package main

import "fmt"

func main() {
    var mat [4][4]int
    for f :=0 ; f < 4; f++ {
        for c := 0; c < 4; c++ {
            fmt.Print("Ingrese componente:")
            fmt.Scan(&mat[f][c])
        }
    }
    fmt.Println("Matriz completa.")
	for i := 0; i < 4; i++{
		
		fmt.Println("")
		fmt.Print("[")
		for p := 0; p < 4; p++ {
            fmt.Print(mat[i][p]," ")
        }
		fmt.Print("]")
	} 
    //fmt.Println(mat)    
    fmt.Println("Elementos de la diagonal principal.")
    for k := 0; k < 4; k++ {
        fmt.Println(mat[k][k]," ")
    }    
}
