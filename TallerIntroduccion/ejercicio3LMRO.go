package main

import "fmt"

func main() {
	var valor1, valor2, valor3, valor4 int
	var suma, promedio int
	fmt.Print("Ingrese el primer valor:")
	fmt.Scan(&valor1)
	fmt.Print("Ingrese el segundo valor:")
	fmt.Scan(&valor2)
	fmt.Print("Ingrese el tercer valor:")
	fmt.Scan(&valor3)
	fmt.Print("Ingrese el cuarto valor:")
	fmt.Scan(&valor4)
	suma = valor1 + valor2 + valor3 + valor4
	promedio = suma / 4
	fmt.Println("La suma de los cuatro valores es:", suma)
	fmt.Print("El promedio de los cuatro valores es:", promedio)
}
