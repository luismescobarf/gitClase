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