package main

import "fmt"

func cuadrado() {
	var var1, var2, var3 int

	fmt.Print("Ingrese el valor 1: ")
	fmt.Scan(&var1)
	fmt.Print("Ingrese el valor 2: ")
	fmt.Scan(&var2)
	fmt.Print("Ingrese el valor 3: ")
	fmt.Scan(&var3)
	if var1 < var2 && var1 < var3 {
		fmt.Print("el valor menor es  ", var1)
	} else if var2 < var1 && var2 < var3 {
		fmt.Print("el valor menor es  ", var2)
	} else {
		fmt.Print("el valor menor es  ", var3)
	}
}
func main() {

	fmt.Print("Hola: ")
	cuadrado()

}
