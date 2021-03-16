package main

import "fmt"

func main() {
	var precio float32
	var cantidad int
	var importe float32
	fmt.Print("Ingrese el precio del producto:")
	fmt.Scan(&precio)
	fmt.Print("Ingrese la cantidad de producto a llevar:")
	fmt.Scan(&cantidad)
	importe = float32(cantidad) * precio
	fmt.Print("El importe total a pagar es:", importe)
}
