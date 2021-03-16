// Se tienen las notas del primer parcial de los alumnos de dos cursos, el curso A y el curso B, 
//cada curso cuenta con 5 alumnos. Realizar un programa que muestre el curso que obtuvo el mayor promedio general.
package main

import "fmt"

func main() {
    var cursoA [5]float32
    var cursoB [5]float32
    var suma1, suma2 float32
    

    fmt.Println("Digitar las notas del curso A")
    for i := 0; i < len(cursoA); i++ {
        fmt.Print("Ingrese nota (valores enteros):")
        fmt.Scan(&cursoA[i])
        suma1 = suma1 + cursoA[i]
    }

    fmt.Println("\n\nDigitar las notas del curso B")
    for i := 0; i < len(cursoB); i++ {
        fmt.Print("Ingrese nota (valores enteros):")
        fmt.Scan(&cursoB[i])
        suma2 = suma2 + cursoB[i]
    }
    
    promedioA := suma1 / 5
    promedioB := suma2 / 5
    fmt.Println("\nPromedio del curso A:", promedioA)
    fmt.Println("\nPromedio del curso B:", promedioB)
    
    if promedioA == promedioB {
        fmt.Println("\nLos cursos tienen el mismo promedio")
    }else{
        if promedioA>promedioB {
            fmt.Println("\nEl curso A tiene un promedio mayor que el curso B.")
        } else {
            fmt.Println("\nEl curso B tiene un promedio mayor que el curso A.")
        }        
    }
    
}