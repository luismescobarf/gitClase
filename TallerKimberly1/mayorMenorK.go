package  main 
import "fmt"

func main(){
var n1, n2 float32
fmt.Println("Leer dos números y si el primero es mayor sumar y restar, sino multiplicar y dividir. ")

fmt.Println("Ingrese el primer número: ")
fmt.Scan(&n1)
fmt.Println("Ingrese el segundo número: ")
fmt.Scan(&n2)

if n1>n2 {
	suma := n1+n2
	resta := n1-n2	

	fmt.Println("La suma de los números es: ", suma, " La resta de los números es igual a: ", resta)
}else{
	div := n1/n2
	multi := n1*n2

	fmt.Println("La multiplicación de los números es: ", multi," La división de los números es: ",div)
}
}