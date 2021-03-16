//Crear y cargar una matriz de 4 filas por 4 columnas. Imprimir la diagonal principal.
package main
import "fmt"

func main(){
	var m[4][4] int
	var l[4] int
	
	for  i := 0; i < 4; i++{
		for  j := 0; j < 4; j++{
		fmt.Println("Ingrese numeros")
		fmt.Scan(&m[i][j])

			if(i==j){
				l[i]=m[i][j]
			}
		}
	}
	fmt.Println(l)
}