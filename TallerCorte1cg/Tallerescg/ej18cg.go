/*Desarrollar una función que solicite la carga de tres valores y muestre el menor.*/
package main
import "fmt"
func nmenor() int{
	var a[3] int
	var n, aux int
	fmt.Println("Ingrese numeros")
	for  i := 0; i < 3; i++{
		fmt.Scan(&n) 
		a[i]=n
	}
	aux=a[0]
	for i := 0; i < 3; i++{
		if(a[i]<aux){ 
		aux=a[i]
		}
	}
	return aux
}

func main(){
	fmt.Println("El numero menor es:",nmenor())
	
}