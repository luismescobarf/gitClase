//Confeccionar un programa que permita almacenar en dos slices los nombres de productos y sus precios 
//(hacer que en las componentes de los mismos subíndices se guarden datos relacionados). 
//Cuando inicia el programa solicitar la cantidad de productos a cargar. Mostrar los datos de los dos slices.
package main

import "fmt"

func main() {
    tamaño := 0
	fmt.Println("¿Cuantos productos registrara?")
	fmt.Scan(&tamaño)

	//Creacion de slices
	productos := make ([]string, tamaño)
	precios := make([]float32, tamaño)

	for i := 0; i < len(productos); i++{
		fmt.Println("\nIngrese el nombre del producto")
		fmt.Scan(&productos[i])
		fmt.Println("Ingrese el precio del producto")
		fmt.Scan(&precios[i])
	}		

	fmt.Println("Datos de los dos slices")
	for i := 0; i < len(productos); i++{
		fmt.Println("\nProducto: ",productos[i]," Precio: ",precios[i])
	}
}	
