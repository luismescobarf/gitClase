/* Ingresar el nombre de 5 productos en un array y sus respectivos precios en otro
array paralelo de tipo float32. Mostrar cuantos productos tienen un precio mayor al
primer producto ingresado (se debe contar) */
package main
import "fmt"

func main(){
	var p[5] string
	var v[5] float32
	var cont int
	for  i := 0; i < 5; i++{
		fmt.Println("Ingrese producto")
		fmt.Scan(&p[i])
		fmt.Println("Ingrese precio")
		fmt.Scan(&v[i])
	}
	aux:=v[0]
	fmt.Println("Productos que tienen un precio mayor al primero:")
	for  i := 1; i < len(v); i++{
		if(v[i]>aux){
			fmt.Println(p[i]) 
			cont++
		}
	}
	fmt.Println("Cantidad de productos",cont) 
}