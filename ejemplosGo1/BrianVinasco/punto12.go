package main

import "fmt"

func main() {
    var mat [2][5]int
    for c :=0 ; c < 5; c++ {
        for f := 0; f < 2; f++ {
            fmt.Print("Ingrese elemento de la fila ", f, " y columna ", c, ":");
            fmt.Scan(&mat[f][c])
        }
    }
    fmt.Println(mat)    
}

