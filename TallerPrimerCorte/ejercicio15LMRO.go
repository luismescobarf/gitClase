package main

import "fmt"

func main() {
	var slice1 []int
	var valor int
	suma := 0
	for {
		fmt.Print("Ingrese un entero (-1 para finalizar):")
		fmt.Scan(&valor)
		if valor == -1 {
			break
		}
		slice1 = append(slice1, valor)
		suma = suma + valor
	}
	promedio := suma / len(slice1)
	mayorProm := 0
	for x := 0; x < len(slice1); x++ {
		if slice1[x] > promedio {
			mayorProm++
		}
	}
	fmt.Println("El slice completo es")
	fmt.Println(slice1)
	fmt.Println("Promedio:", promedio)
	fmt.Println("La cantidad de valores mayores al promedio son:", mayorProm)
}
