package main

import "fmt"

func main() {
	var nom1, nom2 string

	fmt.Print("Ingrese un nombre ")
	fmt.Scan(&nom1)
	fmt.Print("Ingrese un nombre ")
	fmt.Scan(&nom2)
	if nom1 < nom2 {
		fmt.Print(nom2, nom1)
	} else {
		fmt.Print(nom1, nom2)
	}

}
