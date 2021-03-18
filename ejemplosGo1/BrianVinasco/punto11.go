package main

import "fmt"

func main() {
    var mat [3][4]int
    for f :=0 ; f < 3; f++ {
        for c := 0; c < 4; c++ {
            fmt.Print("Ingrese componente:")
            fmt.Scan(&mat[f][c])
        }
    }
    fmt.Println("Matriz completa")
    fmt.Println(mat)
    fmt.Println("Primer fila de la matriz.")
    for c := 0; c < 4; c++ {
        fmt.Print(mat[0][c]," ")
    }
    fmt.Println()
    fmt.Println("Ultima fila de la matriz")
    for c := 0; c < 4; c++ {
        fmt.Print(mat[2][c]," ")
    }
    fmt.Println()
    fmt.Println("Primer columna de la matriz")
    for f := 0; f < 3; f++ {
        fmt.Print(mat[f][0]," ")
    }
    fmt.Println()
}
