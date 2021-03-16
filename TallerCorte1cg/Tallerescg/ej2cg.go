//Ingresar valores enteros por teclado y sumarlos. Cada vez que se carga 
//un valor pedir si quiere continuar con la carga de otro dato ingresando
//el string "si" o "no".
package main
import "fmt"

func main(){
	var num, sum int
	opc:="si"
	for opc == "si"{ 
	fmt.Println("Ingrese numero")
	fmt.Scan(&num)
	sum+=num
	fmt.Println("¿Desea ingresar mas valores?")
	fmt.Scan(&opc)
	}
	fmt.Println("La suma total de valores cargados",sum)
}