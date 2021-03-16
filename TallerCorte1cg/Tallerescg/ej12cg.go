/*Crear una matriz de 2 filas y 5 columnas. Realizar la carga de componentes por
columna (es decir, primero ingresar toda la primera columna, luego la segunda
columna y así sucesivamente) Imprimir luego la matriz.*/

package main
import "fmt"

func main(){
	var m[2][5] int

	for  j := 0; j < 5; j++{
		for  i := 0; i < 2; i++{
		fmt.Println("Ingrese numeros")
		fmt.Scan(&m[i][j])
		}
	}
	fmt.Println(m)
}