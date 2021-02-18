package main

import "fmt"
func main (){

	var precio float32
	var cantidad float32
	fmt.Print("Ingrese el precio del articulo: ")
	fmt.Scan(&precio)
	fmt.Print("Ingrese la cantidad: ")
	fmt.Scan(&cantidad)
	total :=precio*cantidad
	fmt.Println("El total a pagar es: ", total)
	
}