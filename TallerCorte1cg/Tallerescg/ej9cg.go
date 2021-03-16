/*En un curso de 4 alumnos se registraron las notas de sus exámenes y se deben
procesar de acuerdo con lo siguiente:
a) Ingresar Nombre y Nota de cada alumno (almacenar los datos en estructuras
paralelas)
b) Realizar un listado que muestre los nombres, notas y condición del alumno.
En la condición, colocar "Muy Bueno" si la nota es mayor o igual a 8, "Bueno"
si la nota está entre 4 y 7, y colocar "Insuficiente" si la nota es inferior a 4.
c) Imprimir cuantos alumnos tienen la leyenda “Muy Bueno”. */
package main
import "fmt"
func main(){
	var a[4] string
	var n[4] float32
	var cont int
	for  i := 0; i < len(n); i++{
		fmt.Println("Ingrese alumno")
		fmt.Scan(&a[i])
		fmt.Println("Ingrese nota")
		fmt.Scan(&n[i])
	}
	
	for  i := 0; i < len(n); i++{
		if(n[i]>=8){
			fmt.Println("Alumno",a[i],"con nota:",n[i],"condicion: Muy bueno") 
			cont++
		} else if(n[i]>=4 && n[i]<8){
			fmt.Println("Alumno",a[i],"con nota:",n[i],"condicion: Bueno")
		} else if(n[i]<4){
			fmt.Println("Alumno",a[i],"con nota:",n[i],"condicion: Insuficiente")
		}
	}
	fmt.Println("Cantidad de alumnos con muy bien",cont) 
}