/*Definir un slice vacío. En una estructura repetitiva ingresar enteros y añadirlos al
slice. Cuando se ingrese el número -1 finalizar la carga y mostrar cuantos números
ingresados son mayor al promedio de los valores ingresados.*/
package main
import "fmt"

func main(){
var vacio [] int
var opc, sum, cont int

	for opc !=-1{ 
		fmt.Println("Ingrese numeros")
		fmt.Scan(&opc)
		vacio = append(vacio, opc)
		sum+=opc
	}
	
	vacio2 := vacio[:(len(vacio)-1)]
	sum=(sum+1)/(len(vacio2))
	fmt.Println("Promedio", sum)
	
	for _, num := range vacio2 {
        if(num>sum){
			cont++
		}
    }
	fmt.Println("Cantidad de numeros mayores al promedio",cont)
}