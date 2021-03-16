package main

import (
	"fmt"
	"math"
)

func cuadrado() {
	var i float64
	fmt.Print("Ingrese el valor: ")
	fmt.Scan(&i)
	i = math.Pow(i, 2)
	fmt.Print("la raiz de es: ", i)
}
func producto() {
	var valor, valor2 int
	fmt.Print("Ingrese el primer valor: ")
	fmt.Scan(&valor)
	fmt.Print("Ingrese el segundo valor: ")
	fmt.Scan(&valor2)
	suma := valor * valor2
	fmt.Print("El producto de los dos valores es: ", suma)
}

func main() {
	producto()
	fmt.Print("Segunda funcion: ")
	cuadrado()

}
