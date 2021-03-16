package main

import "fmt"

func main() {
    var mat [3][4]int
    for i :=0 ; i < 3; i++ {
        for ii := 0; ii < 4; ii++ {
            fmt.Print("Ingrese numero:")
            fmt.Scan(&mat[i][ii])
        }
    }
    fmt.Println("Matriz completa")
    fmt.Println(mat)
    fmt.Println("Primer fila de la matriz.")
    for ii := 0; ii < 4; ii++ {
        fmt.Print(mat[0][ii]," ")
    }
    fmt.Println()
    fmt.Println("Ultima fila de la matriz")
    for ii := 0; ii < 4; ii++ {
        fmt.Print(mat[2][ii]," ")
    }
    fmt.Println()
    fmt.Println("Primer columna de la matriz")
    for i := 0; i < 3; i++ {
        fmt.Print(mat[i][0]," ")
    }
    fmt.Println()
}