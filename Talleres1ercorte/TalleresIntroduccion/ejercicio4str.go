package main

import "fmt"

func main() {
	var precio, cant float32

	fmt.Println("Ingrese el precio del articulo")
	fmt.Scan(&precio)
	fmt.Println("Ingrese la cantidad del articulo")
	fmt.Scan(&cant)
	final := precio * cant

	fmt.Println("El valor a pagar es de:", final)

}
