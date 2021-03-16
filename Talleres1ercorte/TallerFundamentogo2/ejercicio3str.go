package main

import "fmt"

func main() {

	var arre1 [8]int

	var sum, suma, cant int
	cant = 0

	for f := 0; f < 8; f++ {
		fmt.Print("Ingrese un valor")
		fmt.Scan(&arre1[f])
		sum += arre1[f]
		if arre1[f] > 36 {
			suma += arre1[f]
		}
		if arre1[f] > 50 {
			cant = cant + 1
		}
	}
	fmt.Println("El arreglo es : ", arre1, "La suma de todos es: ", sum, "La suma de los mayores a 36: ", suma, "La cantidad de mayores a 50: ", cant)

}
