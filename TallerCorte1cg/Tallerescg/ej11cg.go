/*Crear y cargar una matriz de 3 filas por 4 columnas. Imprimir la primera fila. Imprimir
la última fila e imprimir la primera columna.*/
package main
import "fmt"

func main(){
	var m[3][4] int
	var p[4] int
	var u[4] int
	var c[3] int
	
	for  i := 0; i < 3; i++{
		for  j := 0; j < 4; j++{
			fmt.Println("Ingrese numeros")
			fmt.Scan(&m[i][j])

			if (i==0){
				p[j]=m[i][j]
			} else if (i==2){
				u[j]=m[i][j]
			} 
			if (j==0) {
				c[i]=m[i][j]
			}
			
		}
	}
	fmt.Println("Primera fila",p)
	fmt.Println("Ultima fila",u)
	fmt.Println("Primera columna",c)
}