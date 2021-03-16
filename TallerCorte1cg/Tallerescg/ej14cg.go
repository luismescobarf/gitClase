/*Definir una matriz de 2 filas y 5 columnas. Realizar su carga e impresión.
Intercambiar los elementos de la primera fila con la segunda y volver a imprimir la
matriz.*/
package main
import "fmt"

func main(){
	var m[2][5] int
	aux:=0
	for  i := 0; i < 2; i++{
		for  j := 0; j < 5; j++{
		fmt.Println("Ingrese numeros")
		fmt.Scan(&m[i][j])
		}
	}
	fmt.Println("Matriz original:",m)

	for  j := 0; j < 5; j++{
		for  i := 0; i < 1; i++{
		aux=m[i][j]
		m[i][j]=m[i+1][j]
		m[i+1][j]=aux
		}
	}
	fmt.Println("Matriz intercambiando filas original:",m)

}