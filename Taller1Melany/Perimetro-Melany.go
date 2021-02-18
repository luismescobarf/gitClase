package main
import "fmt"

func main (){
	var lado int
	fmt.Print("Ingrese el numero del lado del cuadrado: ")
	fmt.Scan(&lado)
	perimetro := lado*4
	fmt.Println("El perimetro del cuadrado es : ", perimetro)
	}