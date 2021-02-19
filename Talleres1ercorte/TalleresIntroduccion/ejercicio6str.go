package main

import "fmt"

func main() {

	var num1, num2, num3 float32
	fmt.Println("Ingrese 3 notas del estudiante recuerde que la nota va de 1 a 10")
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	fmt.Scan(&num3)
	suma := num1 + num2 + num3
	prome := suma / 3
	if prome >= 7 {

		fmt.Println("Promovido", prome)
	} else {
		fmt.Println("Sigue intentando", prome)
	}

}
