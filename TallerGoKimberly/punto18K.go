//Desarrollar una función que solicite la carga de tres valores y muestre el menor.
package main

import "fmt"

func mostrarMenor()  {
	
	var n1, n2, n3 int
	fmt.Println("1. Ingrese un número entero: ")
	fmt.Scan(&n1)
	fmt.Println("\n2. Ingrese un número entero: ")
	fmt.Scan(&n2)
	fmt.Println("\n3. Ingrese un número entero: ")
	fmt.Scan(&n3)

	if n1<n2 && n1<n3 {
		fmt.Println("\nEl número menor es: ",n1)
	}else{
		if n2<n3{
			fmt.Println("\nEl número menor es: ",n2)
		}else{
			fmt.Println("\nEl número menor es: ",n3)
		}
	}
	
}

func main()  {
	mostrarMenor()
}

