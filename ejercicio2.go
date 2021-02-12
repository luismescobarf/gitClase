package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3, num4 int
	fmt.Println("Ingrese el primer  y segundo numero")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)

	fmt.Println("Ingrese el tercero y cuarto  numero")
	fmt.Scanln(&num3)
	fmt.Scanln(&num4)

	suma := (num1 + num2)
	multiplicar := num3 * num4

	fmt.Printf("la suma es %d + %d = %d", num1, num2, suma)
	fmt.Printf(" el producto es %d * %d = %d", num3, num4, multiplicar)
}
