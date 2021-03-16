package main

import "fmt"

func main() {

	var arre1 [5]float32
	var arre2 [5]float32
	var suma, sumb float32

	for f := 0; f < 5; f++ {
		fmt.Print("Ingrese las notas del curso A ")
		fmt.Scan(&arre1[f])
		fmt.Print("Ingreselas notas del curso B ")
		fmt.Scan(&arre2[f])
		suma += arre1[f] / 5
		sumb += arre2[f] / 5
	}
	if suma > sumb {
		fmt.Println(suma, "El gurpo A tiene mejor promedio", arre1)
	} else if sumb > suma {
		fmt.Println(sumb, "El gurpo B tiene mejor promedio", arre2)
	} else {
		fmt.Println("Ambos gurpos tienen el mismo promedio", arre1, arre2, suma, sumb)
	}
}
