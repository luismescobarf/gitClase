package main

import "fmt"

func main() {
	var nombres [4]string
	var notas [4]float32

	for f := 0; f < 4; f++ {
		fmt.Print("Ingrese nombre del alumno:")
		fmt.Scan(&nombres[f])
		fmt.Print("Ingrese su nota:")
		fmt.Scan(&notas[f])
	}
	for f := 0; f < 4; f++ {
		fmt.Print("Nombre: ", nombres[f], " - Nota: ", notas[f])
		if notas[f] >= 8 {
			fmt.Println("  Muy bueno")
		} else if notas[f] > 4 {
			fmt.Println("  Bueno")
		} else {
			fmt.Println("  Insuficiente")
		}
	}
	cant := 0
	for f := 0; f < 4; f++ {
		if notas[f] >= 8 {
			cant++
		}
	}
	fmt.Print("Cantidad de alumnos muy buenos: ", cant)
}
