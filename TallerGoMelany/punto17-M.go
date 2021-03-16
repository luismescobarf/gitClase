package main

import "fmt"

func cuadrado()  {
	var numero, cuadrado int
	fmt.Println("Funcion que calcula el cuadrado de un número ")
	fmt.Println("Ingrese un número entero: ")
	fmt.Scan(&numero)
	cuadrado= numero*numero
	fmt.Println("El cuadrado del número ",numero," es ",cuadrado)

}

func producto()  {
	var numero1, numero2 int
	fmt.Println("\nFuncion que calcula el producto de dos números ")
	fmt.Println("Ingrese un número entero: ")
	fmt.Scan(&numero1)
	fmt.Println("Ingrese un número entero: ")
	fmt.Scan(&numero2)
	producto := numero1*numero2
	fmt.Println("El producto de los números es ",producto)
}

func main(){
	cuadrado()
	producto()
}