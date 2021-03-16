package main

import "fmt"

func main() {
	var productos [5]string
	var precios [5]float32
	for f := 0; f < 5; f++ {
		fmt.Print("Ingrese el nombre del producto:")
		fmt.Scan(&productos[f])
		fmt.Print("Ingrese el precio del producto")
		fmt.Scan(&precios[f])
	}
	cant := 0
	for f := 1; f < 5; f++ {
		if precios[f] > precios[0] {
			cant++
		}
	}
	fmt.Println("La cantidad de productos con un precio mayor al primero ingresado son:", cant)

	for f := 0; f < 5; f++ {
		fmt.Println("Producto: ", productos[f], " - Precio: ", precios[f])
	}
}
