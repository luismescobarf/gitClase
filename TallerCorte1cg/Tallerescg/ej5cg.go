/* Se tienen las notas del primer parcial de los alumnos de dos cursos, el curso A y el
curso B, cada curso cuenta con 5 alumnos. Realizar un programa que muestre el
curso que obtuvo el mayor promedio general. */
package main
import "fmt"

func main(){
	var array1 [5] float32
	var array2 [5] float32
	var prom1, prom2 float32

	for  i := 0; i < len(array1); i++{
		fmt.Println("Curso A:",i+1)
		fmt.Scan(&array1[i])
		prom1+=array1[i]
		fmt.Println("Curso B:",i+1)
		fmt.Scan(&array2[i])
		prom2+=array2[i]
	}
	
	if(prom1>prom2){
		fmt.Println("El curso que obtuvo el mayor promedio general fue: el curso A con promedio",prom1/5)
	} else if (prom1==prom2){ 
		fmt.Println("Los dos cursos tienen el mismo promedio")
	} else{
		fmt.Println("El curso que obtuvo el mayor promedio general fue: el curso B con promedio",prom2/5)
	}
}