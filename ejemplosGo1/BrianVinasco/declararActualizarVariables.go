package main

import "fmt"

func main() {

	var horasTrabajadas int
	var costoHora float32
	var sueldo float32
	//var otraVariable float32

	fmt.Println("Ingresar horas trabajadas por el empleado")
	fmt.Scan(&horasTrabajadas)
	fmt.Println("Ingresar valor hora")
	fmt.Scan(&costoHora)
	sueldo = costoHora * float32(horasTrabajadas)
	fmt.Printf("El salario es %f", sueldo)

}
