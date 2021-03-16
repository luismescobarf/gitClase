/*Crear una matriz de 3x4. Realizar la carga y luego imprimir el elemento mayor.*/

package main
import "fmt"

func main(){
	var m[3][4] int
	aux:= 0
	for  i := 0; i < 3; i++{
		for  j := 0; j < 4; j++{
		fmt.Println("Ingrese numeros")
		fmt.Scan(&m[i][j])

			if(aux<m[i][j]){
				aux=m[i][j]
			}
		}
	}
	fmt.Println("El elemento mayor es:",aux)
}