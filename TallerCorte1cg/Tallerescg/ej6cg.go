/* Cargar un arreglo de 10 elementos y verificar 
posteriormente si el mismo está
ordenado de menor a mayor. */
package main
import "fmt"

func main(){
	var array1 [10] string
	
	for  i := 0; i < len(array1); i++{
	fmt.Println("Ingrese numeros")
	fmt.Scan(&array1[i])
	}
	aux:=array1[0]
	for  i := 1; i < len(array1); i++{
		if array1[i]<aux {
			fmt.Println("No esta ordenado de menor a mayor")
			break
		} else {
			fmt.Println("Esta ordenado de menor a mayor")
			break
		}
		aux=array1[i]
    }
    
}