package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("Ingrese 2 numeros")
	fmt.Scan(&num1)
	fmt.Scan(&num2)

	if num1 > num2 {
		suma := num1 + num2
		resta := num1 - num2
		fmt.Println("la suma de los 2 numeros es", suma, "la diferencia entre ", num1, "y", num2, "=", resta)
	} else {
		multi := num1 * num2
		divi := num1 / num2
		fmt.Println("El producto es", multi, "la division es:", divi)
	}
}
