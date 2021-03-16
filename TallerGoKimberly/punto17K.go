//Desarrollar un programa con dos funciones diferentes al main. La primera solicita el ingreso de un entero, 
//y muestra el cuadrado de dicho valor. La segunda que solicite la carga de dos valores, y muestre su producto. 
//Llamar desde el main a ambas funciones
package main

import "fmt"

func cuadrado(){
	var numero, cuadrado int
	fmt.Println("Funcion que calcula el cuadrado de un número ")
	fmt.Println("Ingrese un número entero: ")
	fmt.Scan(&numero)
	cuadrado= numero*numero
	fmt.Println("El cuadrado del número ",numero," es ",cuadrado)

}

func producto(){
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
