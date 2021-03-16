package main

import "fmt"

func numeros() {
	var cant, canti, vari int

	for f := 1; f != 0; f = vari {
		fmt.Print("Ingrese el valor : ")
		fmt.Scan(&vari)

		if vari < 0 {
			cant = cant + 1
		} else if vari > 0 {
			canti = canti + 1

		}
	}

	fmt.Print("Cantidad de nuemros positivos ", canti, "cantidad de numeros negativos ", cant)

}
func main() {

	fmt.Print("Hola: ")
	numeros()

}
