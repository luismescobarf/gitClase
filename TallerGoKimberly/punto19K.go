//Confeccionar una función que solicite la carga de valores enteros por teclado (se finaliza al ingresar el cero)
//La función retorna la cantidad de valores positivos y negativos ingresados.
package main

import "fmt"

func positivoNegativo() (int, int)  {
	positivos := 0
	negativos := 0
	var numero int

	for{
		fmt.Println("Ingrese número entero: (Para finalizar digite 0)")
		fmt.Scan(&numero)

		if numero==0 {
			break;
		}else{
			if(numero>0){
				positivos++
			}else{
				negativos++
			}
		}
	}
	return positivos, negativos
}

func main()  {
	numerosPositivos, numerosNegativos := positivoNegativo()
	fmt.Println("La cantidad de números positivos es: ",numerosPositivos)
	fmt.Println("\nLa cantidad de números negativos es: ",numerosNegativos)
}
