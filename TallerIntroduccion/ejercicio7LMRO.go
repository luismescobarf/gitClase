package main

import "fmt"

func main() {
	var num int
	fmt.Print("Ingrese un valor entero comprendido entre 1 y 99:")
	fmt.Scan(&num)
	if num < 10 {
		fmt.Print("Tiene un dígito")
	} else {
		fmt.Print("Tiene dos dígitos")
	}
}
