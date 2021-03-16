package main

import "fmt"

func menorValor() {
	var valor1, valor2, valor3 int
	fmt.Print("Ingrese primer valor:")
	fmt.Scan(&valor1)
	fmt.Print("Ingrese segundo valor:")
	fmt.Scan(&valor2)
	fmt.Print("Ingrese tercer valor:")
	fmt.Scan(&valor3)
	if valor1 < valor2 && valor1 < valor3 {
		fmt.Println("Menor valor:", valor1)
	} else {
		if valor2 < valor3 {
			fmt.Println("Menor valor:", valor2)
		} else {
			fmt.Println("Menor valor:", valor3)
		}
	}
}

func main() {
	menorValor()
	menorValor()
	menorValor()
}
