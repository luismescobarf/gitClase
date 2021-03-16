package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("Ingrese primer valor:")
	fmt.Scan(&num1)
	fmt.Print("Ingrese segundo valor:")
	fmt.Scan(&num2)
	if num1 > num2 {
		var suma, diferencia int
		suma = num1 + num2
		diferencia = num1 - num2
		fmt.Println("La suma de los dos valores es ", suma)
		fmt.Println("La diferencia de los dos valores es ", diferencia)
	} else {
		var producto int
		var division float32
		producto = num1 * num2
		division = float32(num1) / float32(num2)
		fmt.Println("El producto de los dos valores es ", producto)
		fmt.Print("La división de los dos valores es ", division)
	}
}
