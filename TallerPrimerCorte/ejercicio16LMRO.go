package main

import "fmt"

func main() {
	var cantidad int
	fmt.Print("Cuantos productos ingresará:")
	fmt.Scan(&cantidad)
	productos := make([]string, cantidad)
	precios := make([]float64, cantidad)

	for f := 0; f < len(productos); f++ {
		fmt.Print("Ingrese nombre del producto:")
		fmt.Scan(&productos[f])
		fmt.Print("Ingrese el precio del producto:")
		fmt.Scan(&precios[f])
	}

	fmt.Println("Listado de todos los datos ingresados")

	for f := 0; f < len(productos); f++ {
		fmt.Println(productos[f], "-", precios[f])
	}
}
