package main

import "fmt"

func main() {

	var arre1 [5]float32
	var prom float32
	var cont, cant int

	for f := 0; f < 5; f++ {
		fmt.Print("Ingrese las alturas ")
		fmt.Scan(&arre1[f])
		prom += arre1[f] / 5

	}
	for f := 0; f < 5; f++ {
		if arre1[f] > prom {
			cont = cont + 1
		}
		if arre1[f] < prom {
			cant = cant + 1
		}
	}

	fmt.Print("arreglo", arre1, "Promedio", prom, "Mas altos", cont, "Mas bajos", cant)
}
