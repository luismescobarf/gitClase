package main

import "fmt"

func calcularCuadrado() {
	var valor int
	fmt.Print("Ingrese un entero:")
	fmt.Scan(&valor)
	cuadrado := valor * valor
	fmt.Println("El cuadrado es ", cuadrado)
}

func calcularProducto() {
	var valor1, valor2 int
	fmt.Print("Ingrese primer valor:")
	fmt.Scan(&valor1)
	fmt.Print("Ingese segundo valor:")
	fmt.Scan(&valor2)
	producto := valor1 * valor2
	fmt.Println("El producto de los valores es:", producto)
}

func main() {
	calcularCuadrado()
	calcularProducto()
}
