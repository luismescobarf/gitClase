package main
import "fmt"

func main (){
	var lado int
	fmt.Println("Leer el lado de un cuadrado y calcular el perímetro. ")
	fmt.Println("Ingrese el valor del lado del cuadrado: ")
	fmt.Scan(&lado)
	perimetro := lado*4
	fmt.Println("El perímetro del cuadrado es : ", perimetro)
	}