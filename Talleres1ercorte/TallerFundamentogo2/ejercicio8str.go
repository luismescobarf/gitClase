package main

import "fmt"

func main() {
	var arre [5]string
	var arre1 [5]float32
	var cont int

	for f := 0; f < 5; f++ {
		fmt.Print("Ingrese los nombres ", f)
		fmt.Scan(&arre[f])
		fmt.Print("Ingrese el precio ", f)
		fmt.Scan(&arre1[f])

	}
	for f := 0; f < 5; f++ {
		if arre1[0] < arre1[f] {
			cont = cont + 1
		}
	}
	fmt.Print("arreglo de nomres", arre, "arreglo de precios", arre1, "cantidad de mayores ", cont)

}
