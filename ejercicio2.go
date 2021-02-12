package main

import "fmt"

func main() {
	var primerNumero, segundoNumero, tercerNumero, cuartoNumero int
	fmt.Println("Ingrese el valor de un lado del primer numero")
	fmt.Scan(&primerNumero)
	fmt.Println("Ingrese el valor de un lado del segundo numero")
	fmt.Scan(&segundoNumero)
	fmt.Println("Ingrese el valor de un lado del tercero numero")
	fmt.Scan(&tercerNumero)
	fmt.Println("Ingrese el valor de un lado del cuarto numero")
	fmt.Scan(&cuartoNumero)

	fmt.Println("la suma de los dos primeros nuemeros es", primerNumero+segundoNumero)
	fmt.Print("el producto del tercer con el cuarto numero es ", tercerNumero, cuartoNumero)
}
