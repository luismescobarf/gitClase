package main

import "fmt"

func main() {

	var arre1 [4]int
	var arre2 [4]int
	var arre3 [4]int

	for f := 0; f < 4; f++ {
		fmt.Print("Ingrese un valor para el arreglo 1 ")
		fmt.Scan(&arre1[f])
		fmt.Print("Ingrese un valor para el arreglo 2 ")
		fmt.Scan(&arre2[f])
		arre3[f] = arre1[f] + arre2[f]
	}
	fmt.Println("El arreglo 1", arre1, "El arreglo 2 ", arre2, "La suma de los dos arreglos", arre3)

}
