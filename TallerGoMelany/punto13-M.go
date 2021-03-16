package main

import "fmt"

func main() {
    var mat [3][4]int
    for i :=0; i < 3; i++ {
        for ii := 0; ii < 4; ii++ {
            fmt.Print("Ingrese el numero:")
            fmt.Scan(&mat[i][ii])
        }
    }
    fmt.Println("Matriz completa")
    fmt.Println(mat)
    mayo := mat[0][0]
    for i := 0; i < 3; i++ {
        for ii := 0; ii < 4; ii++ {
            if mat[i][ii] > mayo {
                mayo = mat[i][ii]
            }
        }
    }
    fmt.Print("El numero mayor de la matriz: ", mayo)    
}
