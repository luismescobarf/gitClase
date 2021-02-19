package main

import "fmt"

func main() {

	var num int
	fmt.Println("Ingrese un numero entre el 1 y 99 ")
	fmt.Scan(&num)

	if num < 10 {

		fmt.Println("El numero ", num, "solo tiene 1 cifra")
	} else if num >= 10 && num < 99 {
		fmt.Println("El numero ", num, "tiene 2 cifra")
	} else {
		fmt.Println("Error numero fuera del limite", num)
	}

}
