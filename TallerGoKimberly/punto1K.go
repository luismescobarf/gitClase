//Ingresar dos nombres de personas y mostrarlos ordenados en forma alfabética.
package main

import "fmt"

func main()  {
	var nombre1, nombre2 string  

	fmt.Println("Ingrese el nombre de la primera persona: ")
	fmt.Scan(&nombre1)
	fmt.Println("Ingrese el nombre de la segunda persona: ")
	fmt.Scan(&nombre2)

	if nombre1 < nombre2 {
		fmt.Println("Los nombres ordenados alfabeticamente son: \n",nombre1, "\n",nombre2)
	}else{
		fmt.Println("Los nombres ordenados alfabeticamente son: \n",nombre2, "\n ",nombre1)
	}
	
}