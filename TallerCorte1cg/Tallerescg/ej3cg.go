/* 3. Desarrollar un programa que permita ingresar un arreglo de 8 elementos, e informe:
• El valor acumulado de todos los elementos del arreglo.
• El valor acumulado de los elementos del arreglo que sean mayores a 36.
• Cantidad de valores mayores a 50. */
package main
import "fmt"

func main(){
	var elem [8] int
	var sum, sum2 int 
	cant := 0
	for  i := 0; i < len(elem); i++{
		fmt.Scan(&elem[i])
	}
	for _, num := range elem {
        sum += num
		if num > 36{
			sum2 += num
		 if num > 50 {
			cant+=1
			}
		}
    }
	fmt.Println("El valor acumulado", sum,"el valor acumulado de los elementos mayores a 36 es", sum2, "Cantidad de elementos mayores a 50 es", cant)
}