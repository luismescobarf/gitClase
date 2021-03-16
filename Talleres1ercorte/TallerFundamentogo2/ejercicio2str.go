package main

import "fmt"

func main() {

	var desci string
	var num, suma int
	desci = "si"

	for desci == "si" {
		fmt.Println("Ingrese un numero")
		fmt.Scan(&num)
		suma = num + suma
		fmt.Println("Ingrese Si si desea continuar o no de lo contrario")
		fmt.Scan(&desci)

	}
	fmt.Println("Suma de numero", suma)
}
