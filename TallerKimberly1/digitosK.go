package  main 
import "fmt"

func main(){
var n1 int
fmt.Println("Se ingresa por teclado un número de 1 a 99 y mostrar por mensaje cuantos digitos tiene")

fmt.Println("Ingrese el número: ")
fmt.Scan(&n1)

if n1>=100 {
	fmt.Println("Este número no cumple con las condiciones del ejercicio.")
}else{
	if n1<100&&n1>=10 {
		fmt.Println("El número tiene dos digitos")
	}else{
		fmt.Println("El número tiene un digito")
	}	
}

}