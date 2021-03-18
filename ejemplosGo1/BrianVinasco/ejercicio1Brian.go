package main

import "fmt"

func main() {
	var lado, perimetro int
	fmt.Println("Ingrese el valor de un lado del cuadrado")
	fmt.Scan(&lado)
	perimetro = lado * 4
	fmt.Println(perimetro)
	fmt.Print("el perimetro del cuadrado es: ", (perimetro))
}
