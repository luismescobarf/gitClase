package main

import "fmt"

func positivosNegativos() (int, int) {
	var valor int
	positivos := 0
	negativos := 0
	for {
		fmt.Print("Ingrese un valor (0 para finalizar):")
		fmt.Scan(&valor)
		if valor == 0 {
			break
		} else {
			if valor > 0 {
				positivos++
			} else {
				negativos++
			}
		}
	}
	return positivos, negativos
}

func main() {
	positivos, negativos := positivosNegativos()
	fmt.Println("La cantidad de valores positivos ingresados son:", positivos)
	fmt.Println("La cantidad de valores negativos ingresados son:", negativos)
}
