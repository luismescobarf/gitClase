package main

import (
	"fmt"
	"sort"
)

func main() {
	var arre1 [10]int
	for f := 0; f < 10; f++ {
		fmt.Print("Ingrese las notas del curso  ")
		fmt.Scan(&arre1[f])

	}
	sort.Ints(arre1[:])
	fmt.Print(arre1, "arreglo ordenado")

}
