package main

import "fmt"

func main() {
	var vec [8]int
	sumaTotal := 0
	sumaMayor36 := 0
	cant := 0
	for f := 0; f < len(vec); f++ {
		fmt.Print("Ingrese elemento:")
		fmt.Scan(&vec[f])
		sumaTotal = sumaTotal + vec[f]
		if vec[f] > 50 {
			cant++
		}
		if vec[f] > 36 {
			sumaMayor36 = sumaMayor36 + vec[f]
		}
	}
	fmt.Println("La suma de los 8 elementos es:", sumaTotal)
	fmt.Println("La suma de los elementos mayores a 36:", sumaMayor36)
	fmt.Println("La cantidad de valores mayores a 50 es:", cant)
}
