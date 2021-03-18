package main

import "fmt"

func main() {
    var vec [10]int
    for f := 0; f < len(vec); f++ {
        fmt.Print("Ingrese elemento:")
        fmt.Scan(&vec[f])
    }
    orden := 1
    for f := 0; f < len(vec) - 1; f++ {
        if vec[f+1] < vec[f] {
            orden=0;
        }
    }
    if orden==1 {
        fmt.Print("Esta ordenado de menor a mayor")
    } else {
        fmt.Print("No esta ordenado de menor a mayor")
    }    
}