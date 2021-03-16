package main

import "fmt"

func main() {

	var arre1 [4]string
	var notas [4]int
	var cont int

	for f := 0; f < 4; f++ {
		fmt.Printf("Ingrese los nombres ")
		fmt.Scan(&arre1[f])
		fmt.Println("Ingrese las notas")
		fmt.Scan(&notas[f])
	}
	for f := 0; f < 4; f++ {
		fmt.Println("------")
		fmt.Println(arre1[f])
		fmt.Println("-")
		fmt.Println(notas[f])
		fmt.Println("-")
		if notas[f] >= 8 {
			fmt.Println("Muy bueno")
			cont = cont + 1

		} else if notas[f] > 4 && notas[f] < 7 {
			fmt.Println("Bueno ")
		} else if notas[f] <= 4 {
			fmt.Println("Insuficiente")

		}
		fmt.Println("------")

	}
	fmt.Print("Numero de alumnos con nota Muy buena ", cont)

}
