package main

import "fmt"

func main() {
	var lado, perimetro int
	fmt.Print("Ingrese el lado del cuadrado:")
	fmt.Scan(&lado)
	perimetro = lado * 4
	fmt.Print("El perímetro del cuadrado es:", perimetro)
}
