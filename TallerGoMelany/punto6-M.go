package main

import "fmt"

func main() {
    var vec [10]int
    for i := 0; i < len(vec); i++ {
        fmt.Print("Ingrese elemento:")
        fmt.Scan(&vec[i])
    }
    orden := 1
    for i := 0; i < len(vec) - 1; i++ {
        if vec[i+1] < vec[i] {
            orden=0;
        }
    }
    if orden==1 {
        fmt.Print("Esta ordenado de menor a mayor ")
    } else {
        fmt.Print("No esta ordenado de menor a mayor ")
    }    
}