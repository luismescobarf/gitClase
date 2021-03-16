package main

import "fmt"

func main() {
    var mat [2][5]int
    for i :=0 ; i < 5; i++ {
        for ii := 0; ii < 2; ii++ {
            fmt.Print("Ingrese elemento de la fila ", ii, " y columna ", i, ":");
            fmt.Scan(&mat[ii][i])
        }
    }
    fmt.Println(mat)    
}