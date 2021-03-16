/*Confeccionar una función que solicite la carga de valores enteros por teclado (se
	finaliza al ingresar el cero)
	La función retorna la cantidad de valores positivos y negativos ingresados.*/
	package main
	import "fmt"
	func cantidad() int{
		var op, p, n int
		op=1
		for (op!= 0){
			fmt.Println("Ingrese numeros")
			fmt.Scan(&op) 
			if(op<0){
				n++
			} else if (op>=0){
				p++
			}
		}
		fmt.Println("Cantidad de numeros positivos",p)
		return n
	}
	
	func main(){
		fmt.Println("Cantidad de numeros negativos",cantidad())	
	}