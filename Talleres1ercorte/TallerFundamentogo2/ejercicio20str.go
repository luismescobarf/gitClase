package main

import "fmt"

func numeros() {
	var arre1[10]int
	var cant,canti,ca int 

	for f := 0; f < 10; f++ {
		fmt.Print("Ingrese un valor para el arreglo  ")
		fmt.Scan(&arre1[f])

		if arre1[f] < 10 {
			cant = cant + 1
		} else if  arre1[f]<=10 ||  arre1[f]<=99  {
			canti = canti + 1

		}else {
			ca=ca+1
		}
	}

	fmt.Print("Cantidad de numero de 1 digito ", cant, " cantidad de numeros 2 digitos ", canti, " cantidad de numeros de mas de 2 digitos ", ca)
	fmt.Print(arre1)
}
func main() {

	fmt.Print("Hola: ")
	numeros()

}