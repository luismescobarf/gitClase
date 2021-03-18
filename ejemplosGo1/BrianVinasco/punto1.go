package main

import "fmt"

func main() {
	var primerNombre, segundoNombre string
	fmt.Println("Ingrese el nombre de la primera persona")
	fmt.Scan(&primerNombre)
	fmt.Println("Ingrese el nombre de la segunda persona")
	fmt.Scan(&segundoNombre)

	if primerNombre < segundoNombre {
		fmt.Println("/nombre ordenados/")
		fmt.Println(primerNombre)
		fmt.Println(segundoNombre)

	}else{
		fmt.Println("/nombre ordenados/")
		fmt.Println(segundoNombre)
		fmt.Println(primerNombre)
	}
}
