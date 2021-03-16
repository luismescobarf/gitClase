/*Desarrollar un programa con dos funciones diferentes al main. La primera solicita el
ingreso de un entero, y muestra el cuadrado de dicho valor. La segunda que solicite
la carga de dos valores, y muestre su product. Llamar desde el main a ambas
funciones.*/

package main
import "fmt"

func prueba1() int {
	var m int
	fmt.Println("Ingrese numero")
	fmt.Scan(&m)
	return m*m
}

func prueba2() int{
	var m, n int
	fmt.Println("Ingrese numeros")
	fmt.Scan(&m)
	fmt.Scan(&n)
	return m*n
}

func main(){
	fmt.Println("Cuadrado",prueba1())
	fmt.Println("Producto",prueba2())
}