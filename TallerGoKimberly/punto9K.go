//En un curso de 4 alumnos se registraron las notas de sus exámenes y se deben procesar de acuerdo con lo siguiente:
// a) Ingresar Nombre y Nota de cada alumno (almacenar los datos en estructuras paralelas)
// b) Realizar un listado que muestre los nombres, notas y condición del alumno. En la condición, colocar "Muy Bueno" si la nota es mayor o igual a 8,
// "Bueno" si la nota está entre 4 y 7, y colocar "Insuficiente" si la nota es inferior a 4.
// c) Imprimir cuantos alumnos tienen la leyenda “Muy Bueno”.

package main

import "fmt"

func main() {
    var nota [4]float32
    var alumno [4]string
	var condicion [4]string
	total := 0
	
    for i := 0 ; i < len(alumno); i++ {
        fmt.Println("\nIngrese el nombre del alumno:")
        fmt.Scan(&alumno[i])
        fmt.Println("Ingrese la nota: ")
        fmt.Scan(&nota[i])
    }

	for i := 0 ; i < len(alumno); i++ {
        if nota[i]>=8 {
			condicion[i]="Muy bueno"
			total++
		}else{
			if nota[i]>=4 && nota[i]<=7 {
				condicion[i]="Bueno"
			}else{
				condicion[i]="Insuficiente"
			}
		}
    }

     fmt.Println("\n\nTotal de alumnos con la leyenda Muy bueno ")
     fmt.Println(total,"\n Los alumnos son: ")

    for i := 0; i < len(alumno); i++ {
        if condicion[i] == "Muy bueno" {
            fmt.Print("\n",alumno[i])
        }
    }

}